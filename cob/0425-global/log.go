package global

// UseLog 함수는 전역 로그를 사용하는 방법을 보여준다.
func UseLog() error {
	if err := Init(); err != nil {
		return err
	}

	// 다른 패키지 안에 있는 경우, 이 값은 global.WithField와 global.Debug를 통해 가져올 수 있다.
	WithField("key", "value").Debug("hello")
	Debug("test")

	return nil
}
