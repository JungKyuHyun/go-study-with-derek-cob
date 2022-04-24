package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// Unwrap 은 wrap한 에러를 분리시킴.
func Unwrap() {

	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Println("wrapped error: ", err)

	// 오류 타입 구분 후 타입에 따른 주석 출력.
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occurred")
	}
}

// StackTrace 오류 스택 출력
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}