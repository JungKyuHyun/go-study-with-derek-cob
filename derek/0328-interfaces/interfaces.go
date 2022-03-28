package interfaces

import (
	"fmt"
	"io"
	"os"
)

func Copy(in io.ReadSeeker, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)
	// MultiWriter : out && os.Stdout 도 데이터를 쓴다 ex) 로그를 파일로도 출력하고 동시에 표준으로도 출력하고 싶은 경우

	// io.Copy(복사객체, 원본객체) but
	if _, err := io.Copy(w, in); err != nil {
		// in에 대량의 데이터가 있는 경우 위험할 수 있기에 err.
		return err
	}

	in.Seek(0, 0)
	// in 파일의 위치를 이동 ex Cursor. Seek(처음 위치, 이동한 위치)
	buf := make([]byte, 64)
	// 64바이트의 buf 만듦.

	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		// CopyBuffer (대상, 복사되는 소스 , 복사 허용 단위) but 버퍼의 길이가 0이면 panic 발생
		return err
	}

	fmt.Println()
	return nil
}
