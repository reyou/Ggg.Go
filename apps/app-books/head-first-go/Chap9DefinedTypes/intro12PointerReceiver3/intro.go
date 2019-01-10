package main

import "fmt"

// MyType qqq
type MyType string

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod()
	pointer.pointerMethod()
}
