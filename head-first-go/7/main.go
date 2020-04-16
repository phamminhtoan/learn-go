package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("%10v | %10v | %10v | %10v | %10v\n", "Type of x", "Type of &x", "&x", "y := &x", "*y")
	fmt.Println("--------------------------------------------------------------------")
	var myInt int = 4
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Printf("%10v | %10v | %10v | %10v | %10v\n", reflect.TypeOf(myInt), reflect.TypeOf(&myInt), &myInt, myIntPointer, *myIntPointer)
	*myIntPointer = 10
	fmt.Printf("%10v | %10v | %10v | %10v | %10v\n", reflect.TypeOf(myInt), reflect.TypeOf(&myInt), &myInt, myIntPointer, *myIntPointer)
	myInt = 20
	fmt.Printf("%10v | %10v | %10v | %10v | %10v\n", reflect.TypeOf(myInt), reflect.TypeOf(&myInt), &myInt, myIntPointer, *myIntPointer)
	var myFloat float64
	myFloatPointer := &myFloat
	fmt.Printf("%10v | %10v | %10v | %10v | %10v\n", reflect.TypeOf(myFloat), reflect.TypeOf(&myFloat), &myFloat, myFloatPointer, *myFloatPointer)
	var myString string
	myStringPointer := &myString
	fmt.Printf("%10v | %10v | %10v | %10v | %10v\n", reflect.TypeOf(myString), reflect.TypeOf(&myString), &myString, myStringPointer, *myFloatPointer)
	var myBool bool
	myBoolPointer := &myBool
	fmt.Printf("%10v | %10v | %10v | %10v | %10v\n", reflect.TypeOf(myBool), reflect.TypeOf(&myBool), &myBool, myBoolPointer, *myBoolPointer)
	ex()
}

func ex() {
	var myInt int = 42
	var myIntPointer *int = &myInt
	fmt.Println(&myInt, myInt, myIntPointer, *myIntPointer, &myIntPointer)
	myInt = 500
	fmt.Println(&myInt, myInt, myIntPointer, *myIntPointer, &myIntPointer)

}
