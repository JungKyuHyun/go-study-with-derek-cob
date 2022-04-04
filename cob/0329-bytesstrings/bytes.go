package bytesstrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer 함수는 Buffer 함수에서 생성한 버퍼를 사용할 것이다.
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"
	b := Buffer(rawString)
	// b.Bytes()를 사용해 버퍼를 바이트로 다시 변환하거나 b.String()를 사용해 버퍼를 문자열로 다시 빠르게 변환할 수 있다.
	fmt.Println(b.String())

	s, err := ToString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// 바이트를 가져와 bytes reader를 생성할 수 있으며
	// 이 reader들은 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, io.RuneScanner 인터페이스를 구현한다.
	reader := bytes.NewBuffer([]byte(rawString))

	// 버퍼링 처리된 읽기 및 토큰화(tokenzation)를 허용하는 스캐너(scanner)에 연결할 수도 있다.
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
