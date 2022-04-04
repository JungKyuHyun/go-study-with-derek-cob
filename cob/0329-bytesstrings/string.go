package bytesstrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString 함수는 문자열을 검색하는 여러 방법을 보여준다.
func SearchString() {
	s := "this is a test"
	fmt.Println(strings.Contains(s, "this")) // true
	// s는 럴자 a를 포함하기 때문에 true를 반환한다.
	fmt.Println(strings.ContainsAny(s, "abc")) // true
	fmt.Println(strings.HasPrefix(s, "this"))  // true
	fmt.Println(strings.HasSuffix(s, "test"))  // true
}

// ModifyString 함수는 문자열을 다양한 방식으로 수정한다.
func ModifyString() {
	s := "simple string"
	fmt.Println(strings.Split(s, " ")) // [simple string]
	fmt.Println(strings.Title(s))      // "Simple String"
	fmt.Println(strings.TrimSpace(s))  // "simple string" (앞뒤 공백 제거)
}

func StringReader() {
	s := "simple string"
	r := strings.NewReader(s)
	// 표준 출력(Stdout)에 출력한다.
	io.Copy(os.Stdout, r)
}
