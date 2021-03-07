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
	"github.com/vtesin/StonkTendency/api/handler"
	"github.com/vtesin/StonkTendency/api/middleware"
	"github.com/vtesin/StonkTendency/config"
	"github.com/vtesin/StonkTendency/infrastructure/repository"
	"github.com/vtesin/StonkTendency/usecase/ticker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Print("1")
	dataSourceName := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dataSourceName))
	fmt.Print("2")
	tickerRepo := repository.NewTickerMongo(client)
	tickerService := ticker.NewService(tickerRepo)

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

	fmt.Print("3")
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.APIPort),
		Handler:      gcontext.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	fmt.Print("4")
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
