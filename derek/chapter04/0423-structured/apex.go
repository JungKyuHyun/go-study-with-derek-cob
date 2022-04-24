package structured

import (
	"errors"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// ThrowError throws error
func ThrowError() error {
	err := errors.New("a crazy failure")
	log.WithField("id", "123").Trace("ThrowError").Stop(&err)
	return err
}

type CustomHandler struct {
	id      string
	handler log.Handler
}

// HandleLog : hook 더해서 로그 기록..
func (h *CustomHandler) HandleLog(e *log.Entry) error {
	e.WithField("id", h.id)
	return h.handler.HandleLog(e)
}

// Apex : 핸들링러 세팅 후 구조화 시켜서 에러 전달
func Apex() {
	log.SetHandler(&CustomHandler{"123", text.New(os.Stdout)})
	err := ThrowError()

	log.WithError(err).Error("an error occurred")
}