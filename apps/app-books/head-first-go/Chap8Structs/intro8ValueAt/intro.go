package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value myStruct
	value.myField = 3
	pointer := &value
	fmt.Println((*pointer).myField)
	fmt.Println(pointer.myField)
	pointer.myField = 10
	fmt.Println((*pointer).myField)
	fmt.Println(pointer.myField)
}
