package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	subscriber1.rate = 4.99
	subscriber1.active = true
	fmt.Println("Name:", subscriber1.name)
	fmt.Println("Montly rate:", subscriber1.rate)
	fmt.Println("Active?", subscriber1.active)
	fmt.Println("///////////////////////////")
	var subscriber2 subscriber
	subscriber2.name = "Beth Ryan"
	subscriber2.rate = 2.99
	subscriber2.active = true
	fmt.Println("Name:", subscriber2.name)
	fmt.Println("Montly rate:", subscriber2.rate)
	fmt.Println("Active?", subscriber2.active)

}
