package main

import (
	"fmt"
	"time"
)

// Pump represents a Gas Pump
type Pump struct {
	Num          int // Represents the pump number
	FillupsGiven int // Represents the number of cars that have been servuces with fuel
}

// fuel is used to provide fuel to a car and increments car.Fillups as well as pump.FillupsGiven
func (pump *Pump) fuel(c chan *Car, q chan int) {
	for {
		select {
		case car := <-c:
			time.Sleep(50 * time.Millisecond)
			car.Fillups++
			pump.FillupsGiven++
		case <-q:
			return
		}
	}
}

// CreatePumps takes in an array int count, and returns an initialized Pump array of size count
func CreatePumps(count int) []Pump {
	pumps := make([]Pump, count)

	// Initialize Pumps and label them starting at 1
	for i := range pumps {
		pumps[i] = Pump{i + 1, 0}
	}

	return pumps
}

// PrintStatus logs the amount of times the pump has given fuel
func (pump *Pump) PrintStatus() {
	fmt.Printf("Pump %v gave fuel %v times\n", pump.Num, pump.FillupsGiven)
}
