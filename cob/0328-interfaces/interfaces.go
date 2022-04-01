package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy 함수는 버퍼를 사용해 in에서 out의 첫 번째 디렉터리로 데이터를 복사한다.
// 데이터를 복사. 표준 출력(stdout)에도 데이터를 쓴다(write).
func Copy(in io.ReadSeeker, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	in.Seek(0, 0)
	// 64바이트 청크(chunk)를 사용해 버퍼에 데이터를 쓴다.
	buf := make([]byte, 64)
	// 표준 복사, 매개변수 in에 대량의 데이터가 있는 경우 이 방법은 위험할 수 있다.ㄴ
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}
	// 새 명령줄에 출력한다.
	fmt.Println()
	return nil
}