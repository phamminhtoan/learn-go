package main

import "fmt"

type Minutes int
type Hours int
type Weight float64
type Title string
type Answer bool

func main() {
	minutes := Minutes(34)
	hours := Hours(2)
	weight := Weight(33.3)
	name := Title("Minh Toan")
	answer := Answer(true)

	minutes += 3
	if weight > 5 {
		fmt.Println(true)
	}
	fmt.Println(minutes, hours, weight, name, answer)
}
