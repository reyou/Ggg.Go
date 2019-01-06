package main

import "fmt"

var myStruct struct {
	number float64
	word   string
	toggle bool
}

func main() {
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Printf("%#v\n", myStruct)
	fmt.Println(myStruct.number)
}
