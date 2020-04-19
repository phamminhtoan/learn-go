package main

import (
	"fmt"
)

type Liters float64
type Milliliters float64
type Gallons float64

func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

type Tilte string

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from ", m)
}

type Number int

func (n *Number) Display() {
	fmt.Println(*n)
}

func (n *Number) Double() {
	*n *= 2
}

func main() {
	var carFuel Gallons
	var busFuel Liters

	carFuel = Gallons(20.0)
	// carFuel := Gallons(20.0)
	busFuel = Liters(240.0) + 50
	// busFuel := Liters(240.0)

	carFuel = Gallons(Liters(30) * 0.264)
	fmt.Println(carFuel, busFuel)

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := Milliliters(500)
	fmt.Printf("%.3f milliters equals %.3f gallons\n", water, water.ToGallons())
	fmt.Printf("%.3f milliters equals %.3f liters\n", water, water.ToLiters())

	milk := Gallons(2)
	fmt.Printf("%.2f gallons equals %.2f liters\n", milk, milk.ToLiters())
	fmt.Printf("%.2f gallons equals %.2f milliliters\n", milk, milk.ToMilliliters())

	fmt.Println(Tilte("Toan") > Tilte("Pham"))

	value := MyType("a MyType value")
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()

	number := Number(4)
	number.Display()
	number.Double()
	number.Display()
}
