package nulls

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type nullInt64 sql.NullInt64

const (
	jsonBlob     = `{"name": "Aaron"}`
	fulljsonBlob = `{"name":"Aaron", "age":0}`
)
// sql 패키지에서 제공하는 null 타입, null 이외의 값을반환하면 유효한 값 설정, 아니면 null 설정함.
type ExampleNullInt struct {
	Age  *nullInt64 `json:"age,omitempty"`
	Name string     `json:"name"`
}

func (v *nullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

func (v *nullInt64) UnmarshalJSON(b []byte) error {
	v.Valid = false
	if b != nil {
		v.Valid = true
		return json.Unmarshal(b, &v.Int64)
	}
	return nil
}

func NullEncoding() error {
	e := ExampleNullInt{}

	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("nullInt64 Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("nullInt64 Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("nullInt64 Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("nullInt64 Marshal, with age = 0:", string(value))

	return nil
}