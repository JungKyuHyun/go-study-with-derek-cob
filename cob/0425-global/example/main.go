package main

import (
	global "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0425-global"
)

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
