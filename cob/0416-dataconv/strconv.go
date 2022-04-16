package dataconv

import (
	"fmt"
	"strconv"
)

// Strconv함수는 일부 strconv함수의 예제를 보여준다.
func Strconv() error {
	// strconv는 문자열에서 다른 타입으로 또는 다른 타입에서 문자열로 변환하는 좋은 방법이다.
	s := "1234"
	// 진수 설정에 10을 지정하고, 정확도 설정에 64비트를 지정할 수 있다.
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	fmt.Println(res)

	// 16진수의 변환
	res, err = strconv.ParseInt("FF", 16, 64)
	if err != nil {
		return err
	}
	fmt.Println(res)

	// bool변환
	val, err := strconv.ParseBool("true")
	if err != nil {
		return err
	}
	fmt.Println(val)

	return nil
}
