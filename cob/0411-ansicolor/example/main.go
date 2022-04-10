package main

import (
	"fmt"

	ansicolor "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0411-ansicolor"
)

func main() {
	r := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text:      "I'm red!",
	}

	fmt.Println(r.String())

	r.TextColor = ansicolor.Green
	r.Text = "Now I'm green!"

	fmt.Println(r.String())

	r.TextColor = ansicolor.ColorNone
	r.Text = "Back to normal..."

	fmt.Println(r.String())
}
