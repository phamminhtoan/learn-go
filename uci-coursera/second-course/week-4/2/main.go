package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct{}

type snake struct{}

type bird struct{}

func (cow) Eat() {
	fmt.Println("grass")
}

func (cow) Move() {
	fmt.Println("walk")
}

func (cow) Speak() {
	fmt.Println("moo")
}

func (bird) Eat() {
	fmt.Println("worms")
}

func (bird) Move() {
	fmt.Println("fly")
}

func (bird) Speak() {
	fmt.Println("peep")
}

func (snake) Eat() {
	fmt.Println("mice")
}

func (snake) Move() {
	fmt.Println("slither")
}

func (snake) Speak() {
	fmt.Println("hiss")
}

func readStrings() []string {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		scannedLn := scanner.Text()
		processedLn := strings.Split(scannedLn, " ")
		if len(processedLn) != 3 {
			fmt.Println("as the input can not be computed please renter the following field")
			continue
		}
		return processedLn
	}
}

func makeStruct(typ string) (Animal, error) {
	switch typ {
	case "cow":
		ret := cow{}
		return ret, nil
	case "bird":
		ret := bird{}
		return ret, nil
	case "snake":
		ret := snake{}
		return ret, nil
	default:
		err := fmt.Errorf("cant make%v", typ)
		return nil, err
	}
}

func MakeAnimal(arr []string, name map[string]Animal) error {
	key := arr[1]
	typ := arr[2]
	strct, err := makeStruct(typ)
	if err == nil {
		name[key] = strct
		fmt.Println("Created It!")
	}
	return err
}

func pullup(arr []string, name map[string]Animal) (Animal, error) {
	pull, find := name[arr[1]]
	if find != true {
		return pull, errors.New("as the respective animal is not currently stored here we are not able to process the following query")
	}
	return pull, nil
}

func read(arr []string, name map[string]Animal) error {
	pull, err := pullup(arr, name)
	if err != nil {
		return err
	}
	switch arr[2] {
	case "eat":
		pull.Eat()
	case "move":
		pull.Move()
	case "speak":
		pull.Speak()
	default:
		fmt.Println("cant understand", arr)
	}
	return err
}

func divide(arr []string, name map[string]Animal) error {
	if arr[0] == "query" {
		err := read(arr, name)
		return err
	}
	if arr[0] == "newanimal" {
		err := MakeAnimal(arr, name)
		return err
	}
	return errors.New("as the repective field dose not meet requirements we are not able to compute it")
}

func makeMap() map[string]Animal {
	mp := make(map[string]Animal)
	return mp
}

func animals(mp map[string]Animal) error {
	arr := readStrings()
	err := divide(arr, mp)
	return err
}

func main() {
	mp := makeMap()
	for {
		err := animals(mp)
		if err != nil {
			fmt.Println(err)
		}
	}
}
