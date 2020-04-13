package main

import (
	"fmt"
	"log"
)

func main() {
	// readFile("./12/data.txt")
	numbers, err := getFloats("./12/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", numbers)

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/sampleCount)

	myString, err := getString("./12/text.txt")
	fmt.Println(myString)
}
