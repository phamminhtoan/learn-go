package main

import "fmt"

func main() {
	var a int
	a = 1
	var b, c int
	b = 2
	c = 3
	var d = 4
	e := 5
	f := a + b + c + d + e
	fmt.Println(a, b, c, d, e, f)
	g, h := 100, 200
	fmt.Println(g, h)
	otherFunc()
}

func otherFunc() {
	var intNum int = 8
	var floatNum float64 = 3.14

	var intNum2 int = int(floatNum)
	var floatNum2 float64 = float64(intNum)

	fmt.Println(intNum2, floatNum2)
	fmt.Println(float64(intNum) > floatNum)
}
