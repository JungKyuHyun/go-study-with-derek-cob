package log

import (
	"log"

	"github.com/pkg/errors"
)

// 원래 에러 반환.
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError : 원래 에러에 주석 레핑해서 반환.
func PassThroughError() error {
	err := OriginalError()
	return errors.Wrap(err, "in passthrougherror")
}

// FinalDestination : 오류를 처리.
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}