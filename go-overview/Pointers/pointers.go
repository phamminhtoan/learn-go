package main

import "fmt"

func main() {
	var aValue float64 = 1.23
	var aPointer *float64 = &aValue

	fmt.Println("aPointer: ", aPointer)
	fmt.Println("*aPointer: ", *aPointer)

	number := 2.6
	halve(number) // pass value of variable
	fmt.Println("number in main: ", number)

	halve2(&number) // pass address of variable
	fmt.Println("number in main: ", number)

}

func halve(number float64) {
	number = number / 2
	fmt.Println("Number in halve: ", number)
}

func halve2(number *float64) {
	*number = *number / 2
	fmt.Println("*number in halve: ", *number)
}
