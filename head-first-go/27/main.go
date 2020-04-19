package main

import (
	"fmt"
	"head-first-go/27/magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "Minh Toan"
	employee.Salary = 6000
	employee.City = "Ho Chi Minh"
	fmt.Println(employee)
	fmt.Println("Name :", employee.Name)
	fmt.Println("City : ", employee.City)
	fmt.Println(employee.Address.City)
	fmt.Println("--------------------")

	address := magazine.Address{City: "Ha Noi"}
	subscriber := magazine.Subscriber{Name: "Ha Le", HomeAdress: address}
	subscriber.HomeAdress.PostalCode = "2000"
	fmt.Println(subscriber)
}
