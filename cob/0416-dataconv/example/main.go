package main

import dataconv "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0416-dataconv"

func main() {
	dataconv.ShowConv()
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Interfaces()
}
