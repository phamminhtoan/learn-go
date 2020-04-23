package main

import (
	"fmt"
	"time"
)

func abc(channel chan string) {
	fmt.Println(2)
	channel <- "a"
	fmt.Println("8888")
	channel <- "b"
	channel <- "c"
	fmt.Println("eeeeee")
}

func def(channel chan string) {
	channel <- "1"
	channel <- "2"
	channel <- "3"
	fmt.Println("zzzz")
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go def(channel2)
	fmt.Println(99999)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println("oooo")

	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	time.Sleep(time.Second * 5)
}
