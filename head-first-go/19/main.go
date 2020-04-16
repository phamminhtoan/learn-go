package main

import (
	"fmt"
	"sort"
)

func main() {
	var myMap map[string]float64
	// The zero value of the map variable itself is nil.
	myMap = make(map[string]float64)
	myMap["Toan"] = 23
	myMap["Thach"] = 25
	fmt.Println(myMap)

	myMap2 := make(map[string]int)
	myMap2["Ha"] = 23
	fmt.Println(myMap2)

	ranks := make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["gold"])

	var names []string
	for name := range ranks {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, ranks[name])
	}

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime)

	mapLiterals := map[string]float64{"a": 1.2, "b": 5.6}
	fmt.Println(mapLiterals)

	numbers := make(map[string]int)
	numbers["i've been assigned"] = 12
	fmt.Printf("%#v\n", numbers["i've been assigned"])      // --> 12
	fmt.Printf("%#v\n", numbers["i haven't been assigned"]) // --> prints the zero value

	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])

	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	_, ok = counters["c"]
	fmt.Println(ok)

}
