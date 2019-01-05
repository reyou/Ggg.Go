package main

import "fmt"

func main() {
	severalInts(1)
	severalInts(1, 2)
	severalInts(1, 2, 3, 4, 5)
	severalStrings("a")
	severalStrings("a", "abc")
	severalStrings("a", "abc", "d", "e")
	mix(1, false, "a", "asdf", "fdsa")
}
func severalInts(numbers ...int) {
	fmt.Println(numbers)
}
func severalStrings(strings ...string) {
	fmt.Println(strings)
}
func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}
