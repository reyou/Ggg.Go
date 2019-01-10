package main

import "fmt"

// Liters qqq
type Liters float64

// Milliliters qqq
type Milliliters float64

// Gallons qqq
type Gallons float64

// ToGallons convert
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

// ToGallons convert
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

// ToLiters qqq
func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

// ToMilliLiters qqq
func (g Gallons) ToMilliLiters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliLiters())
}
