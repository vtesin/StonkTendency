package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vtesin/StonkTendency/usecase/ticker"
)

func getSentiment(tickerService ticker.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Symbol parameter not found in URL"
		vars := mux.Vars(r)
		t, err := tickerService.GetStonk(vars["ticker_id"])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func createSentiment(tickerService ticker.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Symbol parameter not found in Body"

		var input struct {
			Symbol string `json:"symbol"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		err = tickerService.CreateStonk(input.Symbol, 0.0)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

}

//MakeTickerHandlers make url handlers
func MakeTickerHandlers(r *mux.Router, n negroni.Negroni, tickerService ticker.UseCase) {
	r.Handle("/v1/ticker/sentiment/{ticker_id}", n.With(
		negroni.Wrap(getSentiment(tickerService)),
	)).Methods("GET", "OPTIONS").Name("getSentiment")

	r.Handle("/v1/ticker/sentiment", n.With(
		negroni.Wrap(createSentiment(tickerService)),
	)).Methods("POST", "OPTIONS").Name("createSentiment")
}
