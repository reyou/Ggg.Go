package main

import (
	"bufio"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Println("n: ", n, "err: ", err)
	return len(p), nil
}
func main() {
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 3)
	fmt.Println("bw.Buffered(): ", bw.Buffered())
	bw.Write([]byte{'a'})
	fmt.Println("written to bw: a")
	fmt.Println("bw.Buffered(): ", bw.Buffered())
	bw.Write([]byte{'b'})
	fmt.Println("written to bw: b")
	fmt.Println("bw.Buffered(): ", bw.Buffered())
	bw.Write([]byte{'c'})
	fmt.Println("written to bw: c")
	fmt.Println("bw.Buffered(): ", bw.Buffered())
	bw.Write([]byte{'d'})
	fmt.Println("written to bw: d")
	fmt.Println("bw.Buffered(): ", bw.Buffered())
}
