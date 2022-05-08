package main

import (
	"fmt"

	errwrap "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0424-errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.UnWrap()
	errwrap.StackTrace()
}
