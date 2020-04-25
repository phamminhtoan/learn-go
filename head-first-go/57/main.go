package main

import (
	"log"
	"net/http"
	"text/template"
)

type Invoice struct {
	Name    string
	Paid    bool
	Charges []float64
	Total   float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	data := Invoice{
		Name:    "Toan",
		Paid:    true,
		Charges: []float64{12.5, 1.1, 45, 21},
		Total:   67.11,
	}
	html, err := template.ParseFiles("bill.html")
	check(err)
	err = html.Execute(w, data)
	check(err)
}

func main() {
	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
