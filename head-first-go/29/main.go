package main

import (
	"fmt"
	"head-first-go/29/calendar"
	"log"
)

func main() {
	date := calendar.Date{}
	fmt.Println(date)
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
