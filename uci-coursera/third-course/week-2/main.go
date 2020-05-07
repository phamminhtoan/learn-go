package main

import (
	"fmt"
	"time"
)

// it is race condition because the number which is printed depending on which of operations 1 or 2 finish first
// 1: the value of number is being set to 10
// 2: the value of number is being set to 100 by setValue function

func main() {
	fmt.Println("start")
	var number int
	number = 10
	go setValue(&number)

	fmt.Println(number)
	time.Sleep(time.Second * 1)
}
func setValue(number *int) {
	*number = 100
}
