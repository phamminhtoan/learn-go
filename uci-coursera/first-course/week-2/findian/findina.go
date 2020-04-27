package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter a string: ")
	var myString string
	fmt.Scan(&myString)
	if strings.HasPrefix(myString, "i") && strings.HasSuffix(myString, "n") && strings.Contains(myString, "a") {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)

	x := 7
	switch {
	case x > 3:
		fmt.Printf("1")
	case x > 5:
		fmt.Printf("2")
	case x == 7:
		fmt.Printf("3")
	default:
		fmt.Printf("4")
	}

	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)

	// var x int
	// var y *int
	// z := 3
	// y = &z
	// x = &y
}
