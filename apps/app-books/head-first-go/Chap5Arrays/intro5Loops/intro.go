package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}
}
