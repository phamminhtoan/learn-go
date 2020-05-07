package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	animalMap := make(map[string]Animal)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Create: 'newanimal' 'name' 'animal type(cow, bird, snake)'\nquery: 'query' 'name of your animal' 'method'\n")
	for true {
		fmt.Print("> ")
		input, _ := inputReader.ReadString('\n')
		request := strings.SplitN(input, " ", 3)

		command := strings.TrimSpace(request[0])
		animalName := strings.TrimSpace(request[1])
		animalType := strings.TrimSpace(request[2])
		switch true {
		case command == "newanimal":
			animal := createAnimal(animalType)
			if animal != nil {
				animalMap[animalName] = animal
				fmt.Println("Created it!")
			}
		case command == "query":
			queryAnimal(animalName, animalType, animalMap)
		}
	}
}

func queryAnimal(animalName string, infoName string, animalMap map[string]Animal) {
	if animal, ok := animalMap[animalName]; ok {
		switch infoName {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid info request!")
		}
	} else {
		fmt.Println("The animal with the given name is not available")
	}
}

func createAnimal(animalType string) Animal {
	switch true {
	case animalType == "cow":
		return &Cow{AnimalProps{"grass", "walk", "moo"}}
	case animalType == "bird":
		return &Bird{AnimalProps{"worms", "fly", "peep"}}
	case animalType == "snake":
		return &Snake{AnimalProps{"mice", "slither", "hsss"}}
	default:
		fmt.Println("Invalid animal type!")
		return nil
	}
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalProps struct {
	food       string
	locomotion string
	noise      string
}

type Cow struct {
	AnimalProps
}

func (animal *Cow) Eat() {
	fmt.Println(animal.food)
}

func (animal *Cow) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Cow) Speak() {
	fmt.Println(animal.noise)
}

type Bird struct {
	AnimalProps
}

func (animal *Bird) Eat() {
	fmt.Println(animal.food)
}

func (animal *Bird) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Bird) Speak() {
	fmt.Println(animal.noise)
}

type Snake struct {
	AnimalProps
}

func (animal *Snake) Eat() {
	fmt.Println(animal.food)
}

func (animal *Snake) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Snake) Speak() {
	fmt.Println(animal.noise)
}
