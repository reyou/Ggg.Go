package main

import (
	"fmt"
)

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}
func main() {
	myFloatPointer := createPointer()
	fmt.Println(*myFloatPointer)
}
