package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	PhilosophersCount   int = 5
	MaxConcurrentEaters int = 2
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	currentEatersCountChannel := make(chan struct{}, MaxConcurrentEaters)
	chopsticks := make([]*Chopstick, PhilosophersCount)
	for i := 0; i < PhilosophersCount; i++ {
		chopsticks[i] = new(Chopstick)
	}
	wg.Add(PhilosophersCount)
	philosophers := make([]*Philosopher, PhilosophersCount)
	for i := 0; i < PhilosophersCount; i++ {
		philosophers[i] = &Philosopher{
			i + 1,
			chopsticks[i],
			chopsticks[(i+1)%PhilosophersCount],
		}
		go philosophers[i].Eat(&wg, currentEatersCountChannel)
	}
	wg.Wait()
}

func (p Philosopher) Eat(wg *sync.WaitGroup, currentEatersCountChannel chan struct{}) {
	fmt.Print("asdadsa")
	fmt.Println("Philosopher", p.id, "eat()")
	for i := 0; i < 3; i++ {
		currentEatersCountChannel <- struct{}{}
		randBinaryValue := rand.Intn(2)
		if randBinaryValue == 0 {
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()
		} else {
			p.rightChopstick.Lock()
			p.leftChopstick.Lock()
		}
		fmt.Println("			Philosopher", p.id, ": Starting to eat - round", i+1)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("			Philosopher", p.id, ": Finishing eating - round", i+1)

		if randBinaryValue == 0 {
			p.rightChopstick.Unlock()
			p.leftChopstick.Unlock()
		} else {
			p.leftChopstick.Unlock()
			p.rightChopstick.Unlock()
		}
		<-currentEatersCountChannel
	}
	fmt.Println("Philosopher", p.id, "~eat()")
	wg.Done()
}
