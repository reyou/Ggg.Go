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

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

}
