package main

import (
	"fmt"
	"head-first-go/26/magazine"
)

func main() {
	b := magazine.CreadBook("Go Programing")
	fmt.Println(b.Name, b.Author)

	subscriber := magazine.Subscriber{Name: "Toan", Rate: 3.3, Active: true}
	fmt.Printf("%#v\n", subscriber)

	location := coordinates{latitude: 33, longitude: 33.2}
	fmt.Println(location)
}

type coordinates struct {
	latitude  float64
	longitude float64
}
