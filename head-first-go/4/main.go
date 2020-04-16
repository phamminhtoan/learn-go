package main

import "fmt"

func main() {
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("--------------------------------------")
	fmt.Printf("%12s | %2d\n", "Stapms", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
	fmt.Println("--------------------------------------")
	fmt.Printf("%12s | %t\n", "%d", true)
	fmt.Printf("%12s | %v, %#v\n", "%v", "\t\t", "\t")
	fmt.Printf("%12s | %T,  32\n", "%T", 32)
	fmt.Printf("%12s | %%\n", "%%")
	fmt.Println("--------------------------------------")
	fmt.Printf("%%3.2f: %12.2f\n", 2.333333)
	fmt.Printf("%%.3f: %.3f\n", 22.54631)
	var x int = 3
	var y *int = &x
	fmt.Printf("%v\n", &x)
	fmt.Printf("%v\n", y)
}
