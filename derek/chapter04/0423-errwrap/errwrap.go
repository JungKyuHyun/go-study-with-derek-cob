package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// WrappedError 에러와 설명(주석)을 레핑해줌.
func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

type ErrorTyped struct {
	error
}

// Wrap 할때 오류 호출함.
func Wrap() {
	e := errors.New("standard error")

	fmt.Println("Regular Error - ", WrappedError(e))

	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))

	fmt.Println("Nil -", WrappedError(nil))

}