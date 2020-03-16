package main

import "fmt"

func main() {
	beforeLoop := 888
	for i := 1; i <= 5; i++ {
		//inLoop := 999
		fmt.Println(i)
	}
	fmt.Println(beforeLoop) // OK
	// fmt.Println(i)          // Error!
	// fmt.Println(inLoop)     // Error!
}
