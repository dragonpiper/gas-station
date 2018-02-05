package main

import "fmt"

// Car represents a reak worlds Car
type Car struct {
	Num     int // Car number
	Fillups int // Number of times the car has received gas
}

// CreateCars takes in an array int count, and returns an initialized Car array of size count
func CreateCars(count int) []Car {
	cars := make([]Car, count)

	for i := range cars {
		cars[i] = Car{i + 1, 0}
	}

	return cars
}

// PrintStatus logs the amount of times the car has received fuel
func (car *Car) printStatus() {
	fmt.Printf("Car %v was refueld %v times\n", car.Num, car.Fillups)
}
