package main

import "fmt"

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if ok == true {
		if grade < 60 {
			fmt.Printf("%s is failing!\n", name)
		} else {
			fmt.Printf("%s is passing.\n", name)
		}
	} else {
		fmt.Printf("%s is not registered.\n", name)
	}

}
func main() {
	status("Alma")
	status("Carl")
	status("Rohit")
}
