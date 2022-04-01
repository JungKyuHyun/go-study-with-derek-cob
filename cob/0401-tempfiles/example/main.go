package main

import (
	tempfiles "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0401-tempfiles"
)

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}