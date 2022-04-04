package main

import templates "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0402-templates"

func main() {
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}
	if err := templates.InitTemplates(); err != nil {
		panic(err)
	}
	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}
}
