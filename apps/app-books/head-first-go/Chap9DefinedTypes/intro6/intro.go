package main

import "fmt"

// Liters qqq
type Liters float64

// Milliliters qqq
type Milliliters float64

// Gallons qqq
type Gallons float64

// ToGallons convert
func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

// ToLiters convert
func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters \n", busFuel)
}
