package filedirs

import (
	"errors"
	"io"
	"os"
)

// Operate 함수는 파일 및 디렉터리 조작
func Operate() error {
	// /tmp/example 디렉터리를 생성
	err := os.Mkdir("example_dir", os.FileMode(0755));
	if err != nil {
		return err
	}

	// /tmp 디렉터리로 이동
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	// f는 일반적인 파일 객체이며, 여러 인터페이스를 구현
	// 파일을 열때 적절한 비트를 설정하면 reader or writer로 사용 가능
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	value := []byte("hello")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}

	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)
	if err := f.Close(); err != nil {
		return err
	}
	
	if err := os.Chdir(".."); err != nil {
		return err
	}

	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}