package main

import "fmt"

var metersPerLiter float64

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	fmt.Printf("\nArea is %f \n", area)
	fmt.Printf("%.2f liters needed\n", area/metersPerLiter)
	return area / metersPerLiter
}

func main() {
	metersPerLiter = 10
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)
}
