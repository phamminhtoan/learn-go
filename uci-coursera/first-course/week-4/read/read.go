// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// )

// type name struct {
// 	fname string
// 	lname string
// }

// func main() {

// 	var sliceName []name

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Printf("Enter name of the text file: ")
// 	file, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	readFile(file, &sliceName)
// 	for _, n := range sliceName {
// 		fmt.Printf("first name: %s\nlast name: %s\n", n.fname, n.lname)
// 	}

// }

// func check(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func readFile(file string, sliceName *[]name) {
// 	var n name
// 	f, err := os.Open(strings.TrimSpace(file))
// 	check(err)

// 	value := make([]byte, 20)
// 	space := make([]byte, 1)

// 	for {
// 		_, err = f.Read(value)
// 		if err == io.EOF {
// 			break
// 		}
// 		n.fname = strings.TrimSpace(string(value))
// 		_, err = f.Read(space)
// 		if err == io.EOF {
// 			break
// 		}
// 		_, err = f.Read(value)
// 		if err == io.EOF {
// 			break
// 		}
// 		n.lname = strings.TrimSpace(string(value))
// 		*sliceName = append(*sliceName, n)
// 		_, err = f.Read(space)
// 		if err == io.EOF {
// 			break
// 		}
// 	}

// }

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	fmt.Printf("Enter the name of the text file: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName + ".txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	//fmt.Println(file)

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	fullName, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	var people Name
	var peoples []Name

	for _, each := range fullName {
		if len(each[0]) < 20 && len(each[1]) < 20 {
			people.fname = each[0]
			people.lname = each[1]

			peoples = append(peoples, people)
			fmt.Println("First name: " + people.fname + ", Last name: " + people.lname)
		}
	}
}
