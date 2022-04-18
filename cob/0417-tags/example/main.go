package main

import (
	"fmt"

	tags "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0417-tags"
)

func main() {
	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
