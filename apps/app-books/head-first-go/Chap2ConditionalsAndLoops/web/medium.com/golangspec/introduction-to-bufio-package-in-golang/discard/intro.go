package main

import (
	"bufio"
	"fmt"
)

type R struct{}

func (r *R) Read(p []byte) (n int, err error) {
	fmt.Println("Read")
	copy(p, "abcdefghijklmnop")
	return 16, nil
}
func main() {
	r := new(R)
	br := bufio.NewReaderSize(r, 16)
	buf := make([]byte, 4)
	br.Read(buf)
	fmt.Printf("%q\n", buf)
	/*This method throws away n bytes without even returning it.
	If bufio.Reader buffered so far more than or equal to n then
	it doesn’t have to read anything from io.Reader — it simply
	drops first n bytes from the buffer (source code):*/
	br.Discard(4)
	br.Read(buf)
	fmt.Printf("%q\n", buf)
}
