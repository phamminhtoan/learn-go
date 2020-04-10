package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	amount, err := paintNeeded(-2, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)

	}

	root, err := squareRoot(-9.4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.3f", root)
	}

	quotient, err := divide(3.4, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}

	// dozen := double(4.2)
	// fmt.Println(dozen)

}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / 10, nil
}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}
	return math.Sqrt(number), nil
}

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}

func double(number float64) float64 {
	return number * 2
}
