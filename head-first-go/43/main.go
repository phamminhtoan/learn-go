package main

import "fmt"

func calmDown() {
	fmt.Println(recover())
	// recover()
}

func freakOut() {
	defer calmDown()
	panic("Oh no")
	fmt.Println("I won't be run!")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
