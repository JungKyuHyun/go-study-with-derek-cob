package log

import (
	"log"

	"github.com/pkg/errors"
)

// OriginalError 함수는 원래의 오류를 반환한다.
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError 함수는 OriginalError 함술를 호출하고 래핑을 한 다음 오류를 전달한다.
func PassThroughError() error {
	err := OriginalError()
	// nil에도 동작하기 때문에 오류를 확인할 필요가 없다.
	return errors.Wrap(err, "in passthrougherror")
}

// FinalDestination 함수는 오류를 처리하며 전달은 하지 않는다.
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// 예상치 못한 오류가 발생해 로그를 기록한다.
		log.Panicf("an error occurred: %s\n", err.Error())
		return
	}
}
