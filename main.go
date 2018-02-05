package main

import (
	"time"
)

// start begins the loop of fueling Cars at Pumps for until a timeout of 30 seconds is reached
func start() {
	cars := CreateCars(10)  // array containg initialized Car types
	pumps := CreatePumps(4) // array containg initialized Pump types

	c := make(chan *Car) // channel used to send cars to Pumps
	q := make(chan int)  // channgel used to instruct pumps to stop receiving cars

	// Start new coroutine for each pump for servicing cocurrently
	for i := range pumps {
		pump := &pumps[i]
		go pump.fuel(c, q)
	}

	// timeout channel used to shut down pumps after 30 seconds
	timeout := time.After(time.Second * 30)

	// looping for sending cars to pumps untill timeout is reached
	i := 0
	for {
		select {
		case <-timeout:
			for range pumps {
				q <- -1
			}
			for _, pump := range pumps {
				pump.printStatus()
			}
			for _, car := range cars {
				car.printStatus()
			}
			return
		case c <- &cars[i%len(cars)]:
			i++
		}
	}
}

// main is the entry point for our gas station
func main() {
	start()
}
