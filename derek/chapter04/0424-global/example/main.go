package main

import global "github.com/jungkyuhyun/go-study-with-derek-cob/derek/chapter04/0424-global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}