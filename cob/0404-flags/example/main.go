package main

import (
	"flag"
	"fmt"

	flags "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0404-flags"
)

func main() {
	// 설정을 초기화
	c := flags.Config{}
	c.Setup()

	flag.Parse()
	fmt.Println(c.GetMessage())
}
