package main

import (
	"fmt"
)

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}
func main() {
	myBool := true
	printPointer(&myBool)
}
