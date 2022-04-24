package structured

import "github.com/sirupsen/logrus"

type Hook struct {
	id string
}

// Fire : 로그 기록할때 호출됨.
func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	return nil
}

// Levels : logrus 레벨을 보여줌 .
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Logrus : 몇가지 기본적 내장 함수 실행. 
func Logrus() {
	//SetFormatter 로그 포멧을 세팅해줌.
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "Just now"}

	x := logrus.WithFields(fields)
	x.Warn("warning!")
	x.Error("error!")
}