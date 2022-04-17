package main

import (
	"fmt"

	math "github.com/jungkyuhyun/go-study-with-derek-cob/derek/chapter03/0416-math"
)

func main() {
	math.Examples()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()
}