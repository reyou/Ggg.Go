package main

import "fmt"

// Number qqq
type Number int

// Display qqq
func (n *Number) Display() {
	fmt.Println(*n)
}

// Double qqq
func (n *Number) Double() {
	*n *= 2
}
func main() {
	number := Number(4)
	number.Double()
	number.Display()
}
