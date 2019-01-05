package main

import "fmt"

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	i, j := 1, 4
	slice2 := underlyingArray[i:j]
	fmt.Println(slice2)
	fmt.Println()
	slice3 := underlyingArray[2:5]
	fmt.Println(slice3)
	fmt.Println()
	slice4 := underlyingArray[:3]
	fmt.Println(slice4)
	fmt.Println()
	slice5 := underlyingArray[1:]
	fmt.Println(slice5)
}
