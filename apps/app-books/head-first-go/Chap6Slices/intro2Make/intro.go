package main

import "fmt"

func main() {
	printNotes()
	printPrimes()
}
func printPrimes() {
	fmt.Println("\nprintPrimes:")
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
}
func printNotes() {
	fmt.Println("\nprintNotes:")
	notes := make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes)
}
