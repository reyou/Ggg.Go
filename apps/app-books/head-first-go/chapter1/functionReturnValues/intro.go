package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	math.Floor(2.75)
	strings.Title("head first go")
	fmt.Println("this is done.")
	fmt.Println("math.Floor(2.75):", math.Floor(2.75))
	fmt.Println("math.Ceil(2.75):", math.Ceil(2.75))
	fmt.Println("1==2:", 1 == 2)
	fmt.Println("reflect.TypeOf(42):", reflect.TypeOf(42))
	fmt.Println("reflect.TypeOf(3.14):", reflect.TypeOf(3.14))
}
