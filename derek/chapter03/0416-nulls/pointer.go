package nulls

import (
	"encoding/json"
	"fmt"
)

type ExamplePointer struct {
	// int 포인터 타입
	Age  *int   `json:"age,omitempty"`
	Name string `json:"name"`
}

const (
	jsonBlob     = `{"name": "Aaron"}`
	fulljsonBlob = `{"name":"Aaron", "age":0}`
)

// methods for dealing with nil/omitted values
func PointerEncoding() error {

	e := ExamplePointer{}
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	// 값이 구조체인 경우, %+v 변형은 구조체의 필드명까지 포함
	fmt.Printf("Pointer Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with age = 0:", string(value))

	return nil
}