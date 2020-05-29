package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Api struct {
	ID     string
	Width  int
	Height int
	Urls   Urls
}

type Urls struct {
	Raw     string
	Full    string
	Regular string
	Small   string
	Thumb   string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	url := "https://api.unsplash.com/photos/?client_id=DMPE0VjO6IQ9gRd5r46I1o0nDXrr-BJduD-uXTi5cSQ"
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	// fmt.Println(string(body))
	var api []Api
	jsonErr := json.Unmarshal(body, &api)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for _, item := range api {
		fmt.Println(item.Width, item.Height)
	}
	ts, err := template.ParseFiles("index.html")
	check(err)
	ts.Execute(w, api)

}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", mux)
	check(err)

}
