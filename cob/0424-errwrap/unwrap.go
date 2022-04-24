package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

func UnWrap() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")
	fmt.Println("wrapped error: ", err)

	// 다양한 오류 타입을 처리할 수 있다
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occurred")
	}
}

// StackTrace 함수는 오류에 대한 모든 스택을 출력한다.
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}
