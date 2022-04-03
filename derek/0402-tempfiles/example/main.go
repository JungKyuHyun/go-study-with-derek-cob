package main

import tempfiles "github.com/jungkyuhyun/go-study-with-derek-cob/derek/0402-tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
