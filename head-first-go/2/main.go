package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./2/my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Name(), fileInfo.Size())

}
