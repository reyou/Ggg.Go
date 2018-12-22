package main

import (
	"fmt"
)

func main() {
	var quantity = 4
	var length, width float64 = 1.2, 2.4
	var customerName = "alice"

	fmt.Println("222:", quantity)
	fmt.Println("333, 444:", length, width)
	fmt.Println("555:", customerName)
	fmt.Println("666:", length*width)
}
