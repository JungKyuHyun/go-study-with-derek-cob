package encoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type pos struct {
	X      int
	Y      int
	Object string
}

// GobExample함수는 gob 패키지를 사용하는 방법을 보여준다.
func GobExample() error {
	buffer := bytes.Buffer{}

	p := pos{
		X:      10,
		Y:      15,
		Object: "wrench",
	}

	// p가 인터페이스인 경우, 먼저 gob.Register를 호출해야 한다!

	e := gob.NewEncoder(&buffer)
	if err := e.Encode(&p); err != nil {
		return err
	}

	// 버퍼에 있는 데이터가 바이너리 포맷이므로 출력이 잘 되지 않을 수 있다.
	fmt.Println("Gob Encoded valued length: ", len(buffer.Bytes()))

	p2 := pos{}
	d := gob.NewDecoder(&buffer)
	if err := d.Decode(&p2); err != nil {
		return err
	}

	fmt.Println("Gob Decode value: ", p2)

	return nil
}
