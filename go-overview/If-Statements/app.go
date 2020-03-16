package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("true") // printed
	}
	if false {
		fmt.Println("false") // not printed
	}
	if 1 < 2 {
		fmt.Println("1 < 2") // printed
	}
	if 1 > 2 {
		fmt.Println("1 > 2") // not printed
	}

	// Boolean operators
	if !true {
		fmt.Println("!true") // not printed
	}
	if !false {
		fmt.Println("!false") // printed
	}
	if true && false {
		fmt.Println("true && false") // not printed
	}
	if true && true {
		fmt.Println("true && true") // printed
	}
	if true || false {
		fmt.Println("true || false") // printed
	}
	if true || true {
		fmt.Println("true || true") // printed
	}

	// IF ELSE
	if false {
		fmt.Println("if") // not printed
	} else {
		fmt.Println("else") // printed
	}

	// IF ELSE IF
	if false {
		fmt.Println("if") // not printed
	} else if true {
		fmt.Println("else if") // printed
	} else {
		fmt.Println("else") // not printed
	}
}
