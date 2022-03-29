package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

func WorkWrithBuffer() error {
	rawString := "Derek is Good"
	b := Buffer(rawString)
	fmt.Println(b.String())

	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)
	// string to bytes 로
	reader := bytes.NewReader([]byte(rawString))
	// 스케너 설정
	scanner := bufio.NewScanner(reader)
	// 단어별로 split
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return nil

}
