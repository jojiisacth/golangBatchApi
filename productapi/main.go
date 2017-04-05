package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

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

	products := make([]product, 6)

	products[0] = product{"sweet mangoes", "1222", "mangoes of the time"}
	products[1] = product{"greate grapes", "23455", "grapes of the time"}
	products[2] = product{"greate grapes", "5523455", "grapes of the time"}
	products[3] = product{"Tesco Gala Apple Minimum 5 Pack", "213441", "Tesco Gala Apple Minimum 5 Pack"}
	products[4] = product{"Tesco Braeburn Apple Minimum 5 Pack 670G", "2344333", ""}
	products[5] = product{"Tesco Apple Juice 1 Litre", "0987654444", ""}

	//fmt.Fprintf(w, products)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(products)
	return

}

type product struct {
	Name        string `json:"pname"`
	Id          string `json:"id"`
	Description string
}
