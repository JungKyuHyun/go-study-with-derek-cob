package main

import (
	"fmt"

	nulls "github.com/jungkyuhyun/go-study-with-derek-cob/derek/chapter03/0416-nulls"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.NullEncoding(); err != nil {
		panic(err)
	}
}