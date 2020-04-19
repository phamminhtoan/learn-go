package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 6.0
	myStruct.word = "Toan"
	myStruct.toggle = true
	fmt.Printf("%#v\n", myStruct)

	var subscriber1 subscriber
	subscriber1.name = "toan"
	fmt.Println("Name: ", subscriber1.name)

	var subscriber2 subscriber
	subscriber2.name = "ha"
	fmt.Println("Name: ", subscriber2.name)
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}
