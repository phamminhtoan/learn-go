package main

import "fmt"

func main() {
	status("Toan")
	status("Thach")
	status("Ha")
	ex()
}

func status(name string) {
	var grade float64
	var ok bool
	grades := map[string]float64{"Toan": 0, "Thach": 89}
	grade, ok = grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}

}

func ex() {
	data := []string{"a", "b", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}
	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}
}
