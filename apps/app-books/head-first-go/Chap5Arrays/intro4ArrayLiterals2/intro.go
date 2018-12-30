package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes[0], primes[2], primes[4])
}
