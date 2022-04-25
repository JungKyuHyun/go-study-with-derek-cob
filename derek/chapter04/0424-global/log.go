package global

// UseLog : 전역 로그 설정.
func UseLog() error {
	// 초기화
	if err := Init(); err != nil {
		return err
	}

	// global 로그에 필드와 디버그 함수를 추가해줌.
	WithField("key", "value").Debug("hello")
	Debug("test")

	return nil
}