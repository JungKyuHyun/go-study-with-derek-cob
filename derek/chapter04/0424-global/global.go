package global

import (
	"errors"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log     *logrus.Logger
	initLog sync.Once
)

// Init : 처음 설정하고 자주 호출되면 에러를 발송.
func Init() error {
	err := errors.New("already initialized")
	initLog.Do(func() {
		err = nil
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
	})
	return err
}

// sets the log
func SetLog(l *logrus.Logger) {
	log = l
}

// WithField : 전역로그에 연결된 필드 로그를 내보낸다.
func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

// Debg정보를 전역로그에 연결시킴.
func Debug(args ...interface{}) {
	log.Debug(args...)
}