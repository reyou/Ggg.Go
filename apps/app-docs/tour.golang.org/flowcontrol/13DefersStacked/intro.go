/*Deferred function calls are pushed onto a stack.
When a function returns, its deferred calls are executed
in last-in-first-out order.
To learn more about defer statements
read this blog post.*/
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		fmt.Println("Adding into stack:", i)
		defer fmt.Println("Processing:", i)
	}

	fmt.Println("done")
}
