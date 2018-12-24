package main

import (
	"bufio"
	"errors"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Println("n: ", n, "err: ", err)
	fmt.Printf("Write: %q\n", p)
	return 0, errors.New("boom! - buffer overflow!")
}
func main() {
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	fmt.Println("written to bw: a")
	bw.Write([]byte{'b'})
	fmt.Println("written to bw: b")
	bw.Write([]byte{'c'})
	fmt.Println("written to bw: c")
	bw.Write([]byte{'d'})
	fmt.Println("written to bw: d")
	err := bw.Flush()
	fmt.Println("bw flushed")
	fmt.Println(err)
}
