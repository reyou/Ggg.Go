package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	maxTry := 3
	for x := 0; x <= maxTry; x++ {
		fmt.Print("Make a guess: ")
		input, error := reader.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}
		input = strings.TrimSpace(input)
		guess, error := strconv.Atoi(input)
		if error != nil {
			log.Fatal(error)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Congrats! You won!")
			break
		}
		fmt.Println("You have", maxTry-x, "remaining tries.")
		if maxTry-x == 0 {
			fmt.Println("Sorry! The game is over!")
			fmt.Println("Target was", target)
		}
	}

}
