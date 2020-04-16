package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := "data.txt"
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("os.Open() failed with %s\n", err)
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("ioutil.ReadAll() failed with '%s'\n", err)
	}

	lines := bytes.Split(d, []byte{'\n'})
	fmt.Println(lines)
	fmt.Printf("File %s had %d lines\n", path, len(lines))

}
