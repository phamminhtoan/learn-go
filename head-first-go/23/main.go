package main

import "fmt"

func main() {
	var subscriber struct {
		name   string
		rate   float64
		active bool
	}

	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true

	fmt.Println("Name: ", subscriber.name)
	fmt.Println("Monthly rate: ", subscriber.rate)
	fmt.Println("Active? ", subscriber.active)

	var pet struct {
		name string
		age  int
	}
	pet.name = "Max"
	pet.age = 5
	fmt.Println(pet)

	var pet2 struct {
		name string
		age  int
	}

	pet2.name = "Min"
	pet2.age = 3
	fmt.Println(pet2)
}
