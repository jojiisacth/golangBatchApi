package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"strings"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	router.HandleFunc("/batch/", processBatch)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, wellcome to batch api")

}

func processBatch(w http.ResponseWriter, r *http.Request) {
	batchRequest := BatchRequest{}
	parseData(r, &batchRequest)

	results := getBatchResult(&batchRequest)

	for _, result := range results {
		if result != nil && result.response != nil {
			//fmt.Fprintf(w, "{\"status\":\"%s\",\n\"result\":},", result.response.Status)
			fmt.Fprintf(w, "\n")
			result.response.Write(w)
		}
		if result.err != nil {
			fmt.Fprintf(w, "Error:%s", result.err.Error())
			//fmt.Fprintf(w, result.err.Error())

		}
	}

}

func getBatchResult(batchRequest *BatchRequest) []*HttpResponse {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}

	client := http.Client{}
	for _, rq := range batchRequest.Data {

		switch {
		case strings.EqualFold(rq.Method, "get"):
			go func(url string) {
				fmt.Printf("Fetching %s \n", url)
				resp, err := client.Get(url)
				ch <- &HttpResponse{url, resp, err}
				if err != nil && resp != nil && resp.StatusCode == http.StatusOK {
					resp.Body.Close()
				}
			}(rq.Url)
		case strings.EqualFold(rq.Method, "post"):
			go func(rq Request) {
				fmt.Printf("Posting %s \n", rq.Url)
				data := url.Values{}

				for _, body := range rq.Body {
					data.Add(body.Name, body.Value)
				}

				//resp, err := client.Post(rq.Url, rq.ContentType, bytes.NewBufferString(data.Encode()))

				client := &http.Client{}
				r, _ := http.NewRequest("POST", rq.Url, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
				r.Header.Add("Content-Type", "application/json")
				r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

				resp, err := client.Do(r)

				ch <- &HttpResponse{rq.Url, resp, err}
				if err != nil && resp != nil && resp.StatusCode == http.StatusOK {
					resp.Body.Close()
				}

			}(rq)
		default:
			fmt.Printf("Un supported method '%s' \n", rq.Method)
		}

	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			if r.err != nil {
				fmt.Println("with an error", r.err)
			}
			responses = append(responses, r)
			if len(responses) == len(batchRequest.Data) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return responses
}

func parseData(request *http.Request, batchRequest *BatchRequest) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&batchRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type BatchRequest struct {
	Data []Request
}

type Request struct {
	Url         string
	Method      string
	Body        []Body
	ContentType string
}
type Body struct {
	Name  string
	Value string
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}
