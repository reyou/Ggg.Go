package main

import "fmt"

func main() {
	counter := 0
	for {
		fmt.Println("Keeps writing!", counter)
		counter++
		if counter > 1000 {
			break
		}
	}
}
