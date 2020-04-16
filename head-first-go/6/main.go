package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	err := errors.New("Height can't be negative")
	fmt.Println(err.Error())
	fmt.Println(err)
	// log.Fatal(err)

	err = fmt.Errorf("a height of %.2f is invalid", -3.222222)
	fmt.Println(err.Error())
	fmt.Println(err)

	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)

	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)
}

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

func floatParts(number float64) (intergerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}
