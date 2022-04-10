package main

import (
	"bytes"
	"fmt"
	"interfaces"
)

func main() {
	in := bytes.NewBuffer([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy =")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer =", out.String())
	fmt.Print("stdout on pipexample = ")

	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
