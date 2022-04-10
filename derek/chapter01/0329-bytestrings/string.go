package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// string 검색
func SearchString() {
	s := "test string"
	// test가 안에 있는지? true
	fmt.Println(strings.Contains(s, "test"))
	// s에 abc중 하나라도
	fmt.Println(strings.ContainsAny(s, "abc"))
	// 앞에 시작하면,
	fmt.Println(strings.HasPrefix(s, "string"))
	// 뒤에 있다면,
	fmt.Println(strings.HasSuffix(s, "test"))

}

// string 수정
func ModifyString() {
	s := "burger king"
	// [burger , king]
	fmt.Println(strings.Split(s, " "))
	// Burger King
	fmt.Println(strings.Title(s))

	s = " burger king "
	// 앞뒤 공백 제거
	fmt.Println(strings.TrimSpace(s))
}

// 문자열로  io인터 페이스 생성
func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)

	io.Copy(os.Stdout, r)
}
