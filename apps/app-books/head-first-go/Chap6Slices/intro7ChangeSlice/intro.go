package main

import "fmt"

func main() {
	array2 := [5]string{"a", "b", "c", "d", "e"}
	slice2 := array2[2:5]
	slice2[1] = "X"
	fmt.Println(array2)
	fmt.Println(slice2)
}
