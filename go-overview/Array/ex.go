package main

import "fmt"

func main() {
	var months [3]string
	months[0] = "Apr"
	months[1] = "May"
	months[2] = "Jun"

	salesByMonth := [3]float64{12.3, 52, 11.5}

	for i := 0; i < len(months); i++ {
		fmt.Println(months[i], salesByMonth[i])
	}
	fmt.Println("Other way")
	for i, month := range months {
		fmt.Println(month, salesByMonth[i])
	}
	fmt.Println("Do not you index")
	for _, month := range months {
		fmt.Println(month)
	}
}
