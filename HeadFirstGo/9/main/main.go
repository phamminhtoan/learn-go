package main

import (
	"HeadFirstGo/9/calc"
	"HeadFirstGo/9/dates"
	"HeadFirstGo/9/greeting"
	"fmt"
)

func main() {
	greeting.Hi()
	greeting.Hello()
	fmt.Println(calc.Add(3, 4))
	fmt.Println(calc.Subtract(5, 3))
	fmt.Println(dates.WeeksToDays(2))
}
