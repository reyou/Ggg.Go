package main

import "fmt"

// Liters qqq
type Liters float64
// Gallons qqq
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(Liters(40.0))
	busFuel = Liters(Gallons(63.1))
	fmt.Println(carFuel, busFuel)
}
