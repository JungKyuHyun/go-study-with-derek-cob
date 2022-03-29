package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Buffer(rawString string) *bytes.Buffer {
	// rawString을 바이트로 인코딩 시킴. why??
	rawBytes := []byte(rawString)
	// 방법 1.
	// 객체 생성
	b := new(bytes.Buffer)
	// 해당 rawBytes를 b 버퍼에 작성
	b.Write(rawBytes)
	// 방법 2.
	// 바로 버퍼에 초기화 시킴.
	b = bytes.NewBuffer(rawBytes)
	// 방법 3.
	// 초기에 인코딩 시키며 바로 버퍼에 작성.
	b = bytes.NewBufferString(rawString)

	return b
}

func toString(r io.Reader) (string, error) {
	// ReadAll : 리더 안에 것을 다 읽음. []byte로 리턴.
	b, err := ioutil.ReadAll(r)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
