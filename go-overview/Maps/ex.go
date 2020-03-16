package main

import "fmt"

func main() {
	ages := map[string]float64{}
	ages["Alice"] = 12
	ages["Bob"] = 9
	fmt.Println(ages)
	fmt.Println(ages["Alice"])
	// Other way
	ages2 := map[string]float64{"Toan": 23, "Ha": 23}
	for name, age := range ages2 {
		fmt.Println(name, age)
	}
}
