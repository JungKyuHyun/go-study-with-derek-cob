package filedirs

import (
	"bytes"
	"io"
	"os"
	"strings"
)

// 함수 읽고 복사하는 함수
func Capitalizer(f1 *os.File, f2 *os.File) error {
	// seek : 파일의 position을 나타냄.
	if _, err := f1.Seek(0, io.SeekStart); err != nil {
		return err
	}

	var tmp = new(bytes.Buffer)
	// 버퍼에 f1 내용 담음
	if _, err := io.Copy(tmp, f1); err != nil {
		return err
	}
	// 버퍼 대문자 변환
	s := strings.ToUpper(tmp.String())
	// f2에 s를 복사.
	if _, err := io.Copy(f2, strings.NewReader(s)); err != nil {
		return err
	}
	return nil
}

func CapitalizerExample() error {
	f1, err := os.Create("file1.txt")
	if err != nil {
		return err
	}
	if _, err := f1.Write([]byte(`
    this file contains
    a number of words
    and new lines`)); err != nil {
		return err
	}

	f2, err := os.Create("file2.txt")
	if err != nil {
		return err
	}

	if err := Capitalizer(f1, f2); err != nil {
		return err
	}

	if err := os.Remove("file1.txt"); err != nil {
		return err
	}

	if err := os.Remove("file2.txt"); err != nil {
		return err
	}

	return nil
}
