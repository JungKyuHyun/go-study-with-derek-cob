package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 임시 파일, 디렉토리 작업 하는 기본적인 패턴.
func WorkWithTemp() error {
	// 동일한 파일 저장시, 임시 디렉토리에 저장하는 방식.
	// TempDir 반환할 위치("")에 디렉터리("tmp")를 생성.
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}

	// 함수 종료시 임시 파일의 모든 내용을 삭제.
	defer os.RemoveAll(t)

	// 임시 파일 생성하기 위해 경로가 있는지 체크.
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}
	// 정상 저장시 임시저장 파일 이름 출력.
	fmt.Println(tf.Name())

	return nil
}
