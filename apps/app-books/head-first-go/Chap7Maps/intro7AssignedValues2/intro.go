package main

import "fmt"

func main() {
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
	fmt.Println("///////////////////////////")
	_, ok = counters["a"]
	fmt.Println("a exists?", ok)
	_, ok = counters["d"]
	fmt.Println("d exists?", ok)
}
