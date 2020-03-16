package main

import "fmt"

func main() {
	fmt.Println("You win: ")
	doorNumber := 1
	switch doorNumber {
	case 1:
		fmt.Println("a new car")
		fallthrough
	case 2:
		fmt.Println("a llama")
	case 3:
		fmt.Println("a goat")
	default:
		fmt.Println("nothing")
	}
}
