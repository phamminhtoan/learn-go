// Using Pointer with function
package main

import "fmt"

func main() {
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	var myBool bool = true
	printPointer(&myBool)
	negate(&myBool)
	printPointer(&myBool)

	amount := 6.0
	double(&amount)
	fmt.Println(amount)

}
func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer, myBoolPointer)
}

func double(number *float64) {
	*number *= 2
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
