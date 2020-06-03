package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var x float64 = 3.4
	y := reflect.ValueOf(&x)
	fmt.Println("type of y", y.Type())
	fmt.Println("settability of y", y.CanSet())

	z := y.Elem()
	fmt.Println("type of z", z.Type())
	fmt.Println("settability of z", z.CanSet())
	fmt.Println(z.Interface())
	z.SetFloat(10.0)
	fmt.Printf("x: %f, z: %f",x,z)
}