package handler

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vtesin/StonkTendency/usecase/ticker"
)

func stockSentiment(tickerService ticker.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Symbol parameter not found in URL"
		vars := mux.Vars(r)
		err := tickerService.CreateStonk(vars["ticker_id"], 0.0)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}

//MakeTickerHandlers make url handlers
func MakeTickerHandlers(r *mux.Router, n negroni.Negroni, tickerService ticker.UseCase) {
	r.Handle("/v1/ticker/sentiment/{ticker_id}", n.With(
		negroni.Wrap(stockSentiment(tickerService)),
	)).Methods("GET", "OPTIONS").Name("stockSentiment")
}
