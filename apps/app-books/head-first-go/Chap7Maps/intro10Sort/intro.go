package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	names := []string{}
	gradeList := []float64{}
	for name, grade := range grades {
		names = append(names, name)
		gradeList = append(gradeList, grade)
	}
	sort.Strings(names)
	sort.Float64s(gradeList)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
	for _, grade := range gradeList {
		fmt.Printf("Grade %0.1f%%\n", grade)
	}

}
