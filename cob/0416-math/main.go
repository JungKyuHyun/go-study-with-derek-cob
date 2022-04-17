package math

import (
	"fmt"
	"math"
)

// Examples함수는 math 패키지에서 제공하는 일부 함수의 사용법을 보여준다.
func Examples() {
	// 제곱근
	i := 25
	result := math.Sqrt(float64(i))
	fmt.Println(result)

	// 올림
	result = math.Ceil(9.5)
	fmt.Println(result)

	// 내림
	result = math.Floor(9.5)
	fmt.Println(result)

	// math 패키지는 몇 가지 상수(const)도 제공한다
	fmt.Println("Pi: ", math.Pi, "E: ", math.E)
}
