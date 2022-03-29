package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter1/bytestrings"

func main() {
	err := bytestrings.WorkWrithBuffer()
	if err != nil {
		panic(err)
	}

	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
