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
	names := []string{}
	counts := []int{}
	p("empty names and counts initialized.")
	p("///////////////////////////////////\n")
	counter := 0
	for _, line := range lines {
		counter++
		fmt.Println("main counter:", counter)
		p("checking for main line: " + line)
		matched := false
		fmt.Println("names count:", len(names))
		fmt.Println("counts count:", len(counts))
		icounter := 0
		for i, name := range names {
			icounter++
			fmt.Println("inner counter:", icounter)
			p("checking for line: " + line)
			p("checking for name: " + name)
			if name == line {
				fmt.Println("name matched!")
				counts[i]++
				matched = true
			}
			p("========inner loop===========")
		}
		p("matched: " + strconv.FormatBool(matched))
		if matched == false {
			fmt.Println("Adding new name:", line)
			names = append(names, line)
			counts = append(counts, 1)
		}
		p("********outer loop************")
	}
	fmt.Println("\nPrinting Names")
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
