package main

import "fmt"

func main() {
	fmt.Print("Enter a floating point number: ")
	var myFloat float64
	fmt.Scan(&myFloat)
	myInt := int(myFloat)
	fmt.Println(myInt)

}
