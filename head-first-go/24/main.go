package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func showInfo(p part) {
	fmt.Println("Func showInfo")
	fmt.Println("Description: ", p.description)
	fmt.Println("Count: ", p.count)
}

func minimunOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	var porscher car
	porscher.name = "Porsche 911 R"
	porscher.topSpeed = 323
	fmt.Printf("%#v\n", porscher)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Printf("%#v\n", bolts)

	showInfo(bolts)

	p := minimunOrder("Hello world ")
	showInfo(p)
}
