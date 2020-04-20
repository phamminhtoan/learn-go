package main

import "fmt"

type CoffePot string

func (c CoffePot) String() string {
	return string(c) + " CoffePot"
}

func main() {
	coffePot := CoffePot("LuxBrew")
	fmt.Println(coffePot.String())
	fmt.Print(coffePot, "\n")
	fmt.Println(coffePot)
	fmt.Printf("%s", coffePot)
}
