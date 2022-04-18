package encoding

// gob는 Golang 패키지의 데이터 구조를 직렬화 인코딩 / 디코딩 툴 (RPC 콜에서 주로 사용)
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

func GobExample() error {
	buffer := bytes.Buffer{}
	//네트워크에 전달되는 데이커 케리어 설정.
	p := pos{
		X:      10,
		Y:      15,
		Object: "wrench",
	}

	e := gob.NewEncoder(&buffer)
	if err := e.Encode(&p); err != nil {
		return err
	}
	fmt.Println("Gob Encoded valued length: ", len(buffer.Bytes()))

	p2 := pos{}
	d := gob.NewDecoder(&buffer)
	if err := d.Decode(&p2); err != nil {
		return err
	}

	fmt.Println("Gob Decode value: ", p2)

	return nil
}