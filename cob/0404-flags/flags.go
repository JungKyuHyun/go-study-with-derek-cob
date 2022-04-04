package flags

import (
	"flag"
	"fmt"
)

// Config 구조체는 플래그 값을 저장하는데 사용된다
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// Setup 함수는 전달되는 플래그 값으로 설정 값을 초기화 한다.
func (c *Config) Setup() {
	// 다음과 같이 flag를 직접 설정할 수 있지만,
	// var someVar = flag.String("flag_name", "default_val", "description")
	// 하지만 구조체에 값을 넣는 것이 일반적으로 더 낫다.
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	// 축약어
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")
	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// GetMessage 함수는 모든 내부 설정 변수를 사용해 문장을 반환한다.
func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is NOT awesome"
	}
	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())
	return msg
}
