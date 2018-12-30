package main

import (
	"fmt"
	"log"

	average "./average"
	floats "./datafile"
)

func main() {
	numbers, err := floats.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	average := average.GetAverage(numbers)
	fmt.Printf("Average: %0.2f", average)
}
