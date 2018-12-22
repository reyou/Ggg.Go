package main

// The following simple program demonstrates Go's concurrency
// features to implement an asynchronous program.
// It launches two "goroutines" (lightweight threads): one waits for
// the user to type some text, while the other implements a timeout.
// The select statement waits for either of these goroutines to send a
// message to the main routine, and acts on the first message to arrive
// (example adapted from David Chisnall book).

import (
	"fmt"
	"time"
)

func readword(ch chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word)
	ch <- word
}

func timeout(t chan bool) {
	time.Sleep(5 * time.Second)
	t <- false
}

func main2() {
	t := make(chan bool)
	go timeout(t)

	ch := make(chan string)
	go readword(ch)

	select {
	case word := <-ch:
		fmt.Println("Received", word)
	case <-t:
		fmt.Println("Timeout.")
	}
}
