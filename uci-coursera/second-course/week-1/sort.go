package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {

	fmt.Print("Enter sequence of up to 10 integers in one line: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := strings.Split(scanner.Text(), " ")
    mySlice := createIntSlice(input)
    fmt.Println("Before: ",mySlice)
	bubbleSort(&mySlice)
	fmt.Println("After: ", mySlice)
}

func createIntSlice(nums []string) []int {
    var r []int
    for _, v := range nums {
        i, _ := strconv.Atoi(v)
        r = append(r, i)
    }
    return r
}

func bubbleSort(numbers *[]int){
	len := len(*numbers)
	for i:= 0; i < len; i++ {
		for j:= 0; j < len - i -1; j++{
			if (*numbers)[j] > (*numbers)[j+1]{
				swap(numbers , j)
			}
		}
	}
}


func swap(numbers *[]int, i int) {
	(*numbers)[i], (*numbers)[i+1] = (*numbers)[i+1], (*numbers)[i]
}