package filedirs

import (
	"errors"
	"io"
	"os"
)

func Operate() error {
	// 읽고, 쓰고, 실행할 수 있는 (0755) example_dir 파일 만들기
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		// 파일 생성 실패시 에러
		return err
	}

	// /tmp 이동
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	// 이동한 경로에 파일 만들기. f는 해당 파일의 객체
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	// f에 써줌 count는 write 한 만큼.
	value := []byte("hello\n")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	// write count와 쓸 value 길이가 다르다면,
	if count != len(value) {
		// write에 누락이 생긴것.
		return errors.New("incorrect length returned from write")
	}

	if err := f.Close(); err != nil {
		return err
	}

	// open 시 해당 파일객체 리턴.
	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		return err
	}

	// go to /tmp
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// 파일 경로 초기화, but root 경로 삭제는 위험하기 때문에 사용자 입력으로 해당 기능 실행은 위험함.
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}
