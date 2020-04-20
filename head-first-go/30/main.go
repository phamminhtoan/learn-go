package main

import (
	"fmt"
	"head-first-go/30/geo"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(20)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude(), coordinates.Longitude())
}
