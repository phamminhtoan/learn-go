package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	person := make(map[string]string)
	getInput("name", &person)
	getInput("address", &person)
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(nil)
	}
	fmt.Println("Json binary format: ", jsonData)
	fmt.Println("Json format: ", string(jsonData))
}

func getInput(field string, person *map[string]string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your %s : ", field)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	(*person)[field] = strings.TrimSpace(input)
}
