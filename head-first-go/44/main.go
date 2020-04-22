package main

import "fmt"

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}

}

func main() {
	defer calmDown()
	err := fmt.Errorf("There's an error")
	panic(err)
}
