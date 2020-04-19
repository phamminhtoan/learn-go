package magazine

import "time"

type Book struct {
	Name   string
	Author string
	year   time.Time
}

func CreadBook(name string) *Book {
	var b Book
	b.Name = name
	b.Author = "Toan"
	b.year = time.Now()
	return &b
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}
