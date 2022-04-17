package main

import (
	encoding "github.com/jungkyuhyun/go-study-with-derek-cob/derek/chapter03/0416-encoding"
)

func main() {
	if err := encoding.Base64Example(); err != nil {
		panic(err)
	}

	if err := encoding.Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := encoding.GobExample(); err != nil {
		panic(err)
	}
}