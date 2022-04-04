package main

import (
	"fmt"

	csvformat "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0401-csvformat"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		panic(err)
	}
	if err := csvformat.WriteCSVOutput(); err != nil {
		panic(err)
	}
	buffer, err := csvformat.WriteCSVBuffer()
	if err != nil {
		panic(err)
	}
	fmt.Println("Buffer = ", buffer.String())
}
