package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Call int
}

func (s *SpySleeper) Sleep(){
	s.Call++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep(){
	time.Sleep(1*time.Second)
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep(){
	s.Calls = append(s.Calls, sleep)
}


const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper){
	for i:= countdownStart; i > 0; i--{
		sleeper.Sleep()
		fmt.Fprintln(out , i)
	}
	sleeper.Sleep()
	fmt.Fprint(out,finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}