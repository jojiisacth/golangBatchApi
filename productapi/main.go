package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	router.HandleFunc("/products/", getproducts)

	log.Fatal(http.ListenAndServe(":8081", router))

}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, wellcome to product api")

}

func getproducts(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	sb := string(body)
	log.Println(sb)

}
