package main

import "fmt"

func main() {
	var notes []string
	notes = make([]string, 7)
	notes[0] = "Do"
	notes[1] = "Re"
	notes[2] = "Mi"

	fmt.Println(notes, len(notes))

	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	// primes[7] = 4
	fmt.Println(primes, len(primes))

	olds := []int{1, 3, 5}
	fmt.Println(olds, len(olds))
}
