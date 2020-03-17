package main

import (
	"fmt"
	"strings"
)

type Title string
type MyType float64
type Minutes int

func (t Title) FixCase() string {
	return strings.Title(string(t))
}

func (m MyType) MyMethod() {
	fmt.Println("In My Method on MyType", m)
}

func (m *Minutes) Increment() {
	*m += 1
}

func (m Minutes) print() {
	fmt.Println(m)
}

func main() {
	name := Title("minh toan")
	fixed := name.FixCase()
	fmt.Println(fixed)

	myType := MyType(55.5)
	myType.MyMethod()

	minutes := Minutes(10)
	minutes.print()
	minutes.Increment()
	minutes.print()

	hours := Hours(3)
	hours.Increment()
	fmt.Println(hours)
}

type Hours int

func (h *Hours) Increment() Hours {
	if *h+1 > 23 {
		*h = Hours(0)
		return *h
	}
	*h++
	return Hours(*h)
}
