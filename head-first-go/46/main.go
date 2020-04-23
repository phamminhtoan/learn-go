package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 1; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 1; i < 50; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	time.Sleep(time.Second * 1)
	fmt.Println(" end main()")
}
