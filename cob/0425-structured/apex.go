package structured

import (
	"errors"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// ThrowError 함수는 추적할 오류를 발생시킨다
func ThrowError() error {
	err := errors.New("a crazy failure")
	log.WithField("id", "123").Trace("ThrowError").Stop(&err)
	return err
}

// CustomHandler 구조체는 두 개의 스트림으로 분할시킨다.
type CustomHandler struct {
	id      string
	handler log.Handler
}

// HandleLog 함수는 hook을 더해 로그를 기록한다.
func (h *CustomHandler) HandleLog(e *log.Entry) error {
	e.WithField("id", h.id)
	return h.handler.HandleLog(e)
}

// Apex함수는 몇 가지 유용한 기법을 제공한다.
func Apex() {
	log.SetHandler(&CustomHandler{"123", text.New(os.Stdout)})
	err := ThrowError()

	// 오류 편의 기능 활용
	log.WithError(err).Error("an error occurred")
}
