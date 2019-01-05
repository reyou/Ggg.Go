package main

import "fmt"

func main() {
	numbers := map[string]int{}
	numbers["I've been assigned"] = 12
	fmt.Println(numbers["I've been assigned"])
	fmt.Println(numbers["I've not been assigned"])
	fmt.Println("==========================")
	strings := map[string]string{}
	strings["I've been assigned"] = "hi"
	fmt.Println(strings["I've been assigned"])
	fmt.Println(strings["I've not been assigned"])
	fmt.Println("////////////////////////")
	counters := map[string]int{}
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])
	fmt.Println("********************")
}
