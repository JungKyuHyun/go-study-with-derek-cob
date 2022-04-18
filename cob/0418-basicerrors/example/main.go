package main

import (
	"fmt"

	basicerrors "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0418-basicerrors"
)

func main() {
	basicerrors.BasicErrors()
	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
