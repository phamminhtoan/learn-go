package main

import (
	"fmt"
	"math"
)

func main() {
	sum(1, 2, 3)
	number := []float64{1, 2, 3}
	sum(number...)
	sum()
	maximum(-1, -44, 2, 555, -35)
	inRange(4, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
}

func sum(numbers ...float64) {
	var sum float64 = 0
	for _, number := range numbers {
		// fmt.Println(number)
		sum += float64(number)
	}
	fmt.Printf("Sum %.2f\n", sum)
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	fmt.Printf("maximun of %v: %.2f\n", numbers, max)
	return max
}

func inRange(min float64, max float64, numbers ...float64) {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	fmt.Printf("%v min: %.2f max: %.2f\nresult: %v\n", numbers, min, max, result)
}
