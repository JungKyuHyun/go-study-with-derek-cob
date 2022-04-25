package main

import (
	"fmt"

	errwrap "github.com/jungkyuhyun/go-study-with-derek-cob/derek/chapter04/0423-errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}