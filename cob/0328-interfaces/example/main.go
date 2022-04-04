package main

import (
	"bytes"
	"fmt"

	"github.com/jungkyuhyun/go-study-with-derek-cob/cob/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer =", out.String())
	fmt.Println("stdout on PipeExample =")
	if err := interfaces.PipeExampel(); err != nil {
		panic(err)
	}
}
