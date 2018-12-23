// https://golang.org/pkg/bufio/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input := reader.ReadString('\n')
	fmt.Println(input)
}
