package main

import "fmt"

//Monitor is exported
type Monitor struct {
	Resolution string
	Connector  string
	Value      float64
}

func showFieds(m Monitor) {
	fmt.Println(m.Resolution, m.Connector, m.Value)
}

func main() {
	monitor := Monitor{"1080p", "HDMI", 149}
	monitor2 := Monitor{Resolution: "2K"}
	showFieds(monitor)
	showFieds(monitor2)
	c := Clock{3, 2}
	c = Noon()
	fmt.Println(c)
}

// Clock is exported
type Clock struct {
	Hours   int
	Minutes int
}

// Noon is exported
func Noon() Clock {
	c := Clock{12, 0}
	return c
}
