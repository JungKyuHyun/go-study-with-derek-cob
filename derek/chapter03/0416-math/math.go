package math

import (
	"fmt"
	"math"
)

func Examples() {
	i := 25
	// 제곱근
	result := math.Sqrt(float64(i))
	fmt.Println(result)

	// 올림.
	result = math.Ceil(9.5)
	fmt.Println(result)

	// 버림
	result = math.Floor(9.5)
	fmt.Println(result)

	// 파이나 e 자동으로 가지고 있음.
	fmt.Println("Pi:", math.Pi, "E:", math.E)
}