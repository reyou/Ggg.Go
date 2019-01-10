package main

import "fmt"

// Number qqq
type Number int

// Double qqq
func (n Number) Double() {
	n *= 2
}

func main() {
	number := Number(4)
	fmt.Println("Original value of number:", number)
	number.Double()
	fmt.Println("number after calling Double:", number)
}
