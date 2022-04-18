package nulls

import (
	"encoding/json"
	"fmt"
)

const (
	jsonBlob     = `{"name": "Aaron"}`
	fulljsonBlob = `{"name":"Aaron", "age":0}`
)

type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

//  encoding and decoding normal types
func BaseEncoding() error {
	e := Example{}
	// 첫번째 파라미터에 바이트 array 형태의 데이터, 두번째 파라미터에는 출력되는 구조를 포인터로 지정
	// 리턴값은 에러이고, 에러가 없을 경우, 두번째 파라미터에 원래 데이터 들어간다
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with age = 0:", string(value))

	return nil
}