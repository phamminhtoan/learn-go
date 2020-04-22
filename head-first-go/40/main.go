package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func count(start int, end int) {
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
}

func fib(number int) int {
	if number == 1 || number == 0 {
		return 0
	}
	return number + fib(number-1)
}

func main() {
	// count(1, 5)
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal()
	}
	fmt.Println(fib(n))
}
