package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	x := 0
	e := make(chan int)
	go a(e)
	x = <-e
	fmt.Println(x)
	time.Sleep(time.Second * 1)
}
func a(e chan int) {
	var i int
	for i = 0; i < 5000000; i++ {
	}
	e <- i
}
