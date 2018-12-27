package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("height can't be negative")
	// prints the error message
	fmt.Println(err.Error())
	// prints the error message, then exists the program
	log.Fatal(err)
}
