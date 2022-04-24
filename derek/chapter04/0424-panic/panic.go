package panic

import (
	"fmt"
	"strconv"
)

// Panic : panic()함수는 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴한다.
func Panic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}
	// 0 으로 나눌때 패닋
	a := 1 / zero
	fmt.Println("we'll never get here", a)
}

func Catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred:", r)
		}
	}()
	Panic()
}