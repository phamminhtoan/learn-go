package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter some input: \n")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The input was: %s\n", input)
}
