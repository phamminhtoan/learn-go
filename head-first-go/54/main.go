package main

import (
	"log"
	"net/http"
)

func write(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func frenchHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Salut, web!")
}

func hindiHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Namaste, web!")
}

func vietnamHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Xin chao, web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	http.HandleFunc("/xinchao", vietnamHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
