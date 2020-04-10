package main

import (
	"HeadFirstGo/chapter4/calc"
	"HeadFirstGo/chapter4/dates"
	"HeadFirstGo/chapter4/greeting"
	"fmt"
)

func main() {
	greeting.Hi()
	greeting.Hello()
	fmt.Println(calc.Add(3, 4))
	fmt.Println(calc.Subtract(5, 3))
	fmt.Println(dates.WeeksToDays(2))
}
