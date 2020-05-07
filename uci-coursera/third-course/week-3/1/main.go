package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	c := make(chan []int)

	numbers := GetNumbers()
	fmt.Println("You entered: ", numbers)
	SplitAndSort(numbers, &c)

	merge1 := Merge(<-c, <-c)
	merge2 := Merge(merge1, <-c)
	merge3 := Merge(merge2, <-c)
	fmt.Println("Merged array: ", merge3)
}

func SplitAndSort(numbers []int, c *chan []int) {
	length := len(numbers)
	chunk1 := numbers[:length/4]
	chunk2 := numbers[length/4 : length/2]
	chunk3 := numbers[length/2 : length*3/4]
	chunk4 := numbers[length*3/4:]

	go Sort(1, chunk1, c)
	go Sort(2, chunk2, c)
	go Sort(3, chunk3, c)
	go Sort(4, chunk4, c)
}

func Merge(chunk1 []int, chunk2 []int) []int {
	var array = make([]int, 0)
	pos1 := 0
	pos2 := 0
	maxIndex1 := len(chunk1) - 1
	maxIndex2 := len(chunk2) - 1
	for {
		if pos1 <= maxIndex1 && pos2 <= maxIndex2 {
			if chunk1[pos1] <= chunk2[pos2] {
				array = append(array, chunk1[pos1])
				pos1++
			} else {
				array = append(array, chunk2[pos2])
				pos2++
			}
		} else if pos1 <= maxIndex1 {
			array = append(array, chunk1[pos1])
			pos1++
		} else if pos2 <= maxIndex2 {
			array = append(array, chunk2[pos2])
			pos2++
		} else {
			break
		}
	}
	return array
}

func Sort(chunkNo int, chunk []int, c *chan []int) {
	fmt.Printf("Sorting chunk %d: %v\n", chunkNo, chunk)

	sort.Ints(chunk)
	fmt.Printf("Sorted chunk %d: %v\n", chunkNo, chunk)
	*c <- chunk
}

func GetNumbers() []int {
	numbers := []int{}

	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter an integer or enter 'q' to quit: ")
		scanner.Scan()
		userInput = strings.ToLower(scanner.Text())

		if userInput == "q" {
			break
		}

		number, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Wrong number!")
		} else {
			numbers = append(numbers, number)
		}
	}

	return numbers
}
