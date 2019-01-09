package main

import "fmt"

// MyType qqq
type MyType string

// MethodWithParameters qq
func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

// WithReturn qqq
func (m MyType) WithReturn() int {
	return len(m)
}

func (m MyType) unexportedMethod() {
	fmt.Println("This is not exported.")
}

func main() {
	value := MyType("MyType Value")
	value.MethodWithParameters(4, true)
	fmt.Println(value.WithReturn())
	value.unexportedMethod()
}
