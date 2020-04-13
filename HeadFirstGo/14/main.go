package main

import "fmt"

func main() {
	// test()
	// ex()
	ex2()
}

func ex2() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[1:3]
	slice = append(slice, "x")
	slice = append(slice, "y", "z")
	for _, letter := range slice {
		fmt.Println(letter)
	}
	fmt.Println(array)
}
func ex() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2

	for i, number := range numbers {
		fmt.Println(i, number)
	}

	var letters = []string{
		"a",
		"b",
		"c",
	}

	for i, letter := range letters {
		fmt.Println(i, letter)
	}
}

func test() {
	var notes []string
	notes = make([]string, 7)
	notes[0] = "Do"
	notes[1] = "Re"
	notes[2] = "Mi"

	fmt.Println(notes, len(notes))

	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	primes = append(primes, 7)
	fmt.Println(primes, len(primes))

	olds := []int{1, 3, 5}
	fmt.Println(olds, len(olds))

	myArray := [6]int{1, 2, 3, 4, 5, 6}
	mySlice := myArray[:3]
	mySlice = append(mySlice, 5)
	fmt.Println(myArray, mySlice)

}
