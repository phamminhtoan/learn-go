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
	Number  int    `json:"number"`
	Message string `json:"message"`
	People  []People
}

type People struct {
	Craft string
	Name  string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	url := "http://api.open-notify.org/astros.json"

	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	api := Api{}
	jsonErr := json.Unmarshal(body, &api)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(api)

	ts, err := template.ParseFiles("index.html")
	check(err)
	ts.Execute(w, api.Message)

}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", mux)
	check(err)
}
