package main

import "fmt"

func main() {
	var myslice []string
	fmt.Println(myslice)
	var myarray [5]int
	fmt.Println(myarray)
	var notes []string
	notes = make([]string, 7)
	fmt.Println(notes)
}
