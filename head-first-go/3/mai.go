package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it ?")
	fmt.Println(target)
	success := false
	for i := 10; i > 0; i-- {
		fmt.Printf("You have %d guesses left: ", i)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if num == target {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		} else if num > target {
			fmt.Println("Oops. Your guess was HIGHT.")
		} else {
			fmt.Println("Oops. Your guess was LOW.")
		}

	}
	if !success {
		fmt.Printf("Sorry. You ditn't guess my number. It was: %d", target)
	}

	fmt.Print("Exit")

}
