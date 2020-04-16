package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	useSlice()
	useMap()
}

func useMap() {
	lines, err := getStrings("./18/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	counts := make(map[string]int)
	for _, name := range lines {
		counts[name]++
	}
	// fmt.Println(myMap)
	// name, grade
	for key, value := range counts {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func useSlice() {
	lines, err := getStrings("./18/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
			// fmt.Printf("line: %v, i: %v, name: %v, matched: %v\n", line, i, name, matched)
			// fmt.Printf("%v, %v\n", names, counts)
		}
		if matched == false {
			// fmt.Printf("line: %v, matched: %v\n", line, matched)
			names = append(names, line)
			counts = append(counts, 1)
			// fmt.Printf("%v, %v\n", names, counts)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
func getStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
