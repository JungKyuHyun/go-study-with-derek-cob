package main

import (
	"fmt"

	panic "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0425-panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}