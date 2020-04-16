package main

import (
	"fmt"
	"strings"
)

func main() {
	values := []string{"one", "two", "three"}

	result1 := strings.Join(values, "\n")
	fmt.Println(result1)

	result2 := strings.Join(values, "")
	fmt.Printf("Result with no separator: %s\n", result2)

	values2 = []int{10, 200, 300}

}
