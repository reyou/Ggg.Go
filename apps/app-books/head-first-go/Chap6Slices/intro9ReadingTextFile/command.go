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
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	average := average.GetAverage(numbers)
	fmt.Printf("Sample Count: %0.2f", sampleCount)
	fmt.Println()
	fmt.Printf("Average: %0.2f", average)
}
