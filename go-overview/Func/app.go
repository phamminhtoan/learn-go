package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	fmt.Println(square(3))
	fmt.Println(add(2, 4))
	fmt.Println(subtract(10, 3))

	squareRoot, err := squareRoot(9)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(squareRoot)
}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("can't take square root of negative number")
	}
	return math.Sqrt(x), nil
}

func square(number int) int {
	return number * number
}

func add(a float64, b float64) (sum float64) {
	return a + b
}

func subtract(a, b float64) (difference float64) {
	difference = a - b
	return
}
