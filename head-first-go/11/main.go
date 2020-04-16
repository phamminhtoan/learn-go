package main

import (
	"fmt"
	"time"
)

func main() {
	var myArray [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(myArray)
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	fmt.Println(notes[0])
	fmt.Println(notes[4])

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	primes[2]++
	primes[3]++

	fmt.Println(primes[0], primes[1], primes[2], primes[3], primes[4])

	var dates [3]time.Time
	dates[0] = time.Unix(123456854854, 0)
	dates[1] = time.Now()
	fmt.Println(dates[0], dates[1], dates[2])

	text := [3]string{
		"This is a series of long strings",
		"Which would be awkard to place",
		"together on a sighle line",
	}
	fmt.Println(text)
	fmt.Printf("%#v\n", text)

	for i := 0; i < len(text); i++ {
		fmt.Println(i, text[i])
	}
	fmt.Println("------------------------")
	for index, value := range text {
		fmt.Println(index, value)
	}
	fmt.Println("------------------------")
	for _, value := range text {
		fmt.Println(value)
	}
}
