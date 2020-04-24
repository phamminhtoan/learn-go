package main

import (
	"fmt"
	"strings"
)

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return strings.Join(phrases, " and ")
	} else {
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		result += ", and "
		result += phrases[len(phrases)-1]
		return result
	}
}

func main() {
	list := []string{"apple", "orange"}
	fmt.Println(JoinWithCommas(list))
	list = []string{"apple", "orange", "pear"}
	fmt.Println(JoinWithCommas(list))
}
