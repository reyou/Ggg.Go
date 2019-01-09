package main

import "fmt"

// Liters qqq
type Liters float64

// Gallons qqq
type Gallons float64

func main() {
	 fmt.Println(Liters(1.2) + Liters(3.4))
	 fmt.Println(Gallons(5.5) - Gallons(2.2))
	 fmt.Println(Liters(1.2) / Liters(1.1))
	 fmt.Println(Gallons(1.2) == Gallons(1.2))
	 fmt.Println(Liters(1.2) < Liters(3.4))
	 fmt.Println(Liters(1.2) > Liters(3.4))
}
