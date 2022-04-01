package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WorkWithTemp 함수는 임시 파일 및 디렉터리 작업을 하는 기본적인 패턴을 보여준다.
func WorkWithTemp() error {
	// 예를 들어 template1-10.html과 동일한 이름을 가진 파일을 저장할
	// 임시 공간이 필요한 경우 임시 디렉터리가 좋은 접근 방법이다.
	// 첫번째 인자는 공백인데, 이는 os.TempDir()이 반환한 위치에
	// 디렉터리를 생성한다는 것을 의미한다.
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}
	// 나중에 이 작업을 수행하려는 경우, 함수가 종료될 때 임시 파일의 모든 내용을 삭제한다.
	// 따라서 반드시 디렉터리 이름을 호출 함수에 반환해야 한다. 
	defer os.RemoveAll(t)

	// 임시 파일을 생성하기 위해서는 디렉터리가 반드시 존재해야 한다.
	// t는 *os.File 객체다
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}
	fmt.Println(tf.Name())

	return nil
}