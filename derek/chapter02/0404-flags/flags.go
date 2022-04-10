package main

import (
	"flag"
	"fmt"
)

// 플레그 값을 저장할 Config 구조체
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// 리시버로 Config 의 메소드로 작용.
func (c *Config) Setup() {
	// flag : 커맨드라인 플래그 파싱을 지원하는 패키지임.

	// StringVar : 명령줄 옵션을 받은 뒤 문자열로 저장.
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	// StringVar 축약 버전.
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")
	// 명령줄 옵션을 받은 뒤 bool로 저장
	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	// 명령줄 옵션을 받은 뒤 정수로 저장
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")

	// 명령줄 옵션을 받은 뒤 커스텀으로 저장
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// Config의 속성을 사용해 메시지 출력.
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
