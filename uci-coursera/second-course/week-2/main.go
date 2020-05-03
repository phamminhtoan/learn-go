package main

import (
	"fmt"
)

func main() {
	var a, v, s, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter velocity: ")
	fmt.Scan(&v)
	fmt.Print("Enter displacement: ")
	fmt.Scan(&s)
	fmt.Print("Enter time: ")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(t))
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		result := (a*t*t)/2 + v*t + s
		return result
	}
}
