package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	a[2] = 88
	s1 = s1[0:4]
	fmt.Println(a, s1, s2)

	fmt.Println("len(s1):", len(s1), "cap(s1): ", cap(s1))
	fmt.Println("len(s1):", len(s1), "cap(s1): ", cap(s1))
	s2 = append(s2, 5)

	fmt.Println("len(s1):", len(s1), "cap(s1): ", cap(s1))
	fmt.Println("len(s1):", len(s1), "cap(s1): ", cap(s1))
	fmt.Println(a, s1, s2)

	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println(slice)
	for _, el := range slice {
		fmt.Println(el)
	}
	result, err := First(slice)
	fmt.Println(result, err)
}

func First(slice []int) (int, error) {
	if len(slice) != 0 {
		return slice[0], nil
	}
	return 0, fmt.Errorf("Slice is empty")

}
