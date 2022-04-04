package interfaces

import (
	"io"
	"os"
)

func PipeExampel() error {
	// pipe reader와 pipe writer는 io.Reader와 io.Writer를 구현한다.
	// Pipe()는 동기 인메모리(in-memory) 파이프를 만든다.
	r, w := io.Pipe()

	go func() {
		w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
