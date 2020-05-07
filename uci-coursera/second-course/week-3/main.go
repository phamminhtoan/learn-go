package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AnimalInterFace interface {
	Eat()
	Move()
	Speak()
}

type Animal struct {
	food       string
	locomotion string
	spoken     string
}

func (a Animal) Eat() {
	fmt.Println("result: ", a.food)
}

func (a Animal) Move() {
	fmt.Println("result: ", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println("result: ", a.spoken)
}

func CheckMethod(animal Animal, method string) {
	switch method {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var query, animal, method string
	for {
		getInput(&query, &animal, &method)
		if animal == "cow" {
			CheckMethod(cow, method)
		}
		if animal == "bird" {
			CheckMethod(bird, method)
		}
		if animal == "snake" {
			CheckMethod(snake, method)
		}
	}
}

func getInput(query, animal, method *string) {
	fmt.Println("Please Enter query animal method: ")
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, _ := inputReader.ReadString('\n')
	s := strings.SplitN(input, " ", 3)

	*query = strings.TrimSpace(s[0])
	*animal = strings.TrimSpace(s[1])
	*method = strings.TrimSpace(s[2])

}

func getAnimal() string {
	animalList := []string{"cow", "bird", "snake"}
	var animal string
	for {
		fmt.Print("Enter one of 'cow', 'bird' ,'snake': ")
		fmt.Scan(&animal)
		if contains(animalList, animal) {
			break
		}
	}
	return animal
}

func getMethod() string {
	methodList := []string{"eat", "move", "speak"}
	var method string
	for {
		fmt.Print("Enter one of 'eat', 'move' ,'speak': ")
		fmt.Scan(&method)
		if contains(methodList, method) {
			break
		}
	}
	return method
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
