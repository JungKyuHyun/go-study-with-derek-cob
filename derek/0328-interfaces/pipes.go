package interfaces

import (
	"io"
	"os"
)

func PipeExample() error {
	// Pipe는 새로운 Reader와 Writer를 제공해주거나 Writer를 반전
	r, w := io.Pipe()
	// "go" 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행 = 비동기적 실행 = 동시 실행에 사용.
	// 익명함수임. 마지막 } 뒤에 ()를 사용해 함수 호출. 안에 매개변수 넣을 수 있음.
	go func() {
		w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
