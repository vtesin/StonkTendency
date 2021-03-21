package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	gcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jonreiter/govader"
	"github.com/vtesin/StonkTendency/api/handler"
	"github.com/vtesin/StonkTendency/api/middleware"
	"github.com/vtesin/StonkTendency/config"
	"github.com/vtesin/StonkTendency/infrastructure/repository"
	"github.com/vtesin/StonkTendency/usecase/sentiment"
	"github.com/vtesin/StonkTendency/usecase/ticker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	dataSourceName := fmt.Sprintf("mongodb://%s:%s@%s:%d", config.DbUser, config.DbPassword, config.DbHost, config.DbPort)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dataSourceName))

	if err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err.Error())
		}
	}()

	twitter := sentiment.NewTwitter("https://api.twitter.com/2/tweets/search/recent?query=%23$symbol%20-is:retweet&tweet.fields=created_at,author_id,lang&max_results=100", "$symbol", 100)
	analyzer := govader.NewSentimentIntensityAnalyzer()
	sources := make([]*sentiment.Source, 1)
	sources[0] = &twitter.Source

	tickerRepo := repository.NewTickerMongo(client)
	sentimentService := sentiment.NewService(sources, analyzer)

	tickerService := ticker.NewService(tickerRepo, sentimentService)

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	//ticker
	handler.MakeTickerHandlers(r, *n, tickerService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.APIPort),
		Handler:      gcontext.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
