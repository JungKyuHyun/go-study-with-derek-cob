package panic

import (
	"fmt"
	"strconv"
)

// Panic 함수는 0으로 나누기로 인해 패닉을 발생시킨다
func Panic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}

	a := 1 / zero
	fmt.Println("we'll never get here", a)
}

// Catcher 함수는 Panic함수를 호출한다.
func Catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred:", r)
		}
	}()
	Panic()
}