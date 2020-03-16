package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("Hello, I'm Toan")
	otherFunc()
	roundedUp := math.Ceil(myNumber)
	roundedDown := math.Floor(myNumber)
	fmt.Println(roundedUp, roundedDown)
}

var myNumber = 1.23

func otherFunc() {
	fmt.Println("other function here!")
	fmt.Println("here" + "abc")
	fmt.Println(reflect.TypeOf(1))

}
