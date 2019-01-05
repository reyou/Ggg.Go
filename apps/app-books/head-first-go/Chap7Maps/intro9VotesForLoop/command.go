package main

import (
	"fmt"
	"log"
	"strconv"

	"./datafile"
)

func p(input string) {
	fmt.Println(input)
}
func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := map[string]int{}
	counter := 0
	for _, line := range lines {
		counter++
		fmt.Println("main counter:", counter)
		p("checking for main line: " + line)
		_, ok := counts[line]
		if ok {
			counts[line]++
		} else {
			counts[line] = 1
		}
	}
	fmt.Println("\nPrinting Names")
	for key, value := range counts {
		fmt.Printf("\n%s: %s", key, strconv.Itoa(value))
	}
	fmt.Println("\n\n*******Only Names**********")
	for name := range counts {
		fmt.Println(name)
	}
	fmt.Println("\n\n*******Only Values**********")
	for _, value := range counts {
		fmt.Println(value)
	}
}
