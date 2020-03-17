package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTask(channel chan int) { // Remove return value, accept channel chan int
	delay := rand.Intn(5)
	fmt.Println("Starting long task")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("Long task finished")
	channel <- delay // Changed from "return delay"
}

func main() {
	rand.Seed(time.Now().Unix())
	channel := make(chan int)                 // Add
	go longTask(channel)                      // Change from "time := longTask()"
	fmt.Println("Took", <-channel, "seconds") // "<-channel" instead of "time"
}
