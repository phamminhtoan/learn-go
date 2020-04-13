package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// readFile("./12/data.txt")
	fmt.Println(os.Args[1:])
	numbers, err := getFloats2("./13/data.txt")
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

	myString, err := getString("./13/text.txt")
	fmt.Println(myString)
}
