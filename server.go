package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// POST /buy/ donde el payload es: {'id':1, 'cant':1} y la respuesta es: {'date':'31-07-2020'} dt.now + 5

func main() {

	type Item struct {
		ID       int     `json:"id"`
		Price    float32 `json:"price"`
		Quantity int     `json:"quantity"`
		Date     string  `json:"delivery_date"`
	}

	http.HandleFunc("/buy/", func(w http.ResponseWriter, r *http.Request) {
		item := Item{}
		// 1. llamar a stock service y restarle 1 al id 1
		// 2. desde stock service llamar a delivery service y que segun la cantidad me devuelva la delivery date
		// 3. acá sólo debería devolver la delivery date
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

	log.Print("Listening :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
