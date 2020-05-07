package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	arr := getInput()
	parts := 4

	var wg sync.WaitGroup
	c := make(chan []int, parts)

	n := int(math.Max(math.Ceil(float64(len(arr))/float64(parts)), 1))
	for i := 0; i < parts; i++ {
		start := int(math.Min(float64(i*n), float64(len(arr))))
		end := int(math.Min(float64((i+1)*n), float64(len(arr))))

		wg.Add(1)

		go func(a []int) {
			fmt.Println("sorting: ", a)
			sort.Ints(a)
			c <- a
			wg.Done()
		}(arr[start:end])

	}
	wg.Wait()
	close(c)
	var result []int
	for item := range c {
		for _, z := range item {
			result = append(result, z)
		}
	}
	sort.Ints(result)
	fmt.Println("sorted: ", result)
}

func getInput() []int {
	fmt.Printf("Enter array of integers separated by space :\n")
	r := bufio.NewReader(os.Stdin)
	list, _ := r.ReadString('\n')
	s := strings.Split(strings.TrimSpace(list), " ")
	var arr []int

	for _, str := range s {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, num)
	}
	return arr
}
