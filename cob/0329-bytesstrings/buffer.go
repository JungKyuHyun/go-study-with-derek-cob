package bytesstrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer 함수는 bytes 버퍼를 초기화하는 일부 기법을 보여준다.
// 이 버퍼들은 io.Reader 인터페이스를 구현한다.
func Buffer(rawString string) *bytes.Buffer {
	// 원시 바이트(raw bytes)로 인코딩된 문자열에서 시작. string을 byte 슬라이스로 변환
	// https://yourbasic.org/golang/convert-string-to-byte-slice/
	rawBytes := []byte(rawString)

	/**
	[]byte -> string
	rawBytes := string([]byte(rawString))
	*/

	// 내장 함수 new는 포인터 변수를 생성하며, 제공된 타입의 제로 값을 가리키는 포인터를 반환한다.
	var b = new(bytes.Buffer)

	// 원시 바이트(raw bytes)나 원래의 문자열에서 버퍼를 생성하는 방법 1
	b.Write(rawBytes)
	// 방법 2
	b = bytes.NewBuffer(rawBytes)
	// 방법 3 - 초기 바이트 배열을 완전히 피하는 방법
	b = bytes.NewBufferString(rawString)

	return b
}

// ToString 함수는 io.Reader를 가져와 모두 사용한 다음 문자열을 반환한다.
func ToString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}