package main

import (
	"fmt"
	"head-first-go/9/calc"
	"head-first-go/9/dates"
	"head-first-go/9/greeting"
)

func main() {
	greeting.Hi()
	greeting.Hello()
	fmt.Println(calc.Add(3, 4))
	fmt.Println(calc.Subtract(5, 3))
	fmt.Println(dates.WeeksToDays(2))
}
