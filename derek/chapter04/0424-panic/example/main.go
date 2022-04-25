package main

import (
	"fmt"

	panic "github.com/jungkyuhyun/go-study-with-derek-cob/derek/chapter04/0424-panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}