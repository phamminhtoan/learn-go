package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	ex()
	path := "../data.txt"
	lines, err := ReadLine(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File %s had %d lines\n", path, len(lines))
	fmt.Println(lines)

	content := strings.Join(lines, "\n")
	fmt.Println(content)
}

// Read the whole file
func ex() {
	d, err := ioutil.ReadFile("../data.txt")
	if err != nil {
		log.Fatalf("ioutil.ReadFile failed with %s\n", err)
	}
	fmt.Printf("Size of 'data.txt': %d bytes\n", len(d))
}

// Open file for reading, close file
func ex1() {
	f, err := os.Open("../data.txt")
	if err != nil {
		log.Fatalf("os.File failed with '%s'\n", err)
	}
	defer f.Close()
}

// ReadLine is exported -- Read file line by file
func ReadLine(filePath string) ([]string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
