package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

func main() {
	var result string = stringutil.Reverse("Pham Minh Toan")
	fmt.Println(result)
}
