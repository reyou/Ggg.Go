package main

import (
	"fmt"
)

func main() {
	var quantity int
	var length, width float64
	var customerName string
	quantity = 2
	customerName = "alice"
	length, width = 1.2, 2.4
	fmt.Println("111:", quantity)
	fmt.Println("222, 333:", length, width)
	fmt.Println("444:", customerName)
	fmt.Println("555:", length*width)
}
