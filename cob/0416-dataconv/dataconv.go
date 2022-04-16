package dataconv

import "fmt"

// ShowConv함수는 일부 타입의 변환 예제르 보여준다.
func ShowConv() {
	// int
	var a = 24
	// float64
	var b = 2.0

	// 계산을 위해 int를 float64로 변환
	c := float64(a) * b
	fmt.Println(c)
	// fmt.Sprintf는 문자열로 변환할 떄 많이 사용된다.
	precision := fmt.Sprintf("%.2f", b)

	// 값(%s)과 타입(%T) 출력
	fmt.Printf("%s - %T\n", precision, precision)
}
