package main

import "fmt"

// Liters qqq
type Liters float64

// Gallons qqq
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
}
