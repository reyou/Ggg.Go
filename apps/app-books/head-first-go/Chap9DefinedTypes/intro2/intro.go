package main

import "fmt"

// Liters qqq
type Liters float64
// Gallons qqq
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = 10.0
	busFuel = 240.0
	fmt.Println(carFuel, busFuel)
}
