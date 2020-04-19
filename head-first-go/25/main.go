package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name : ", s.name)
	fmt.Println("Monthly rate: ", s.rate)
	fmt.Println("Active? ", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 2.99
}

func main() {
	subscriber1 := defaultSubscriber("Minh Toan")
	subscriber1.rate = 3
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Minh Thach")
	printInfo(subscriber2)
	applyDiscount(subscriber2)
	printInfo(subscriber2)

	var pointer *subscriber = subscriber2
	fmt.Println(pointer.name)
	(*pointer).name = "Ha Le"
	fmt.Println(pointer.name)
}
