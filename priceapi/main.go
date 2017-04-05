package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	router.HandleFunc("/price/", getPrice)

	log.Fatal(http.ListenAndServe(":8083", router))

}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, wellcome to price api")

}

func getPrice(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	sb := string(body)
	log.Println(sb)

	prices := make([]price, 6)

	prices[0] = price{"1222", "2654", "76", "$"}
	prices[1] = price{"23455", "23455", "76", "$"}
	prices[2] = price{"5523455", "2564", "76", "$"}
	prices[3] = price{"213441", "6542", "76", "$"}
	prices[4] = price{"2344333", "2344", "76", "$"}
	prices[5] = price{"0987654444", "5256", "76", "$"}
	//fmt.Fprintf(w, products)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(prices)
	return
}

type price struct {
	ProductId string `json:"pid"`
	Storeid   string `json:"sid"`
	Value     string
	Currency  string
}
