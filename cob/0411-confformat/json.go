package confformat

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type JSONData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ToJSON 함수는 JSONData 구조체를 YAML 포맷의 bytes.Buffer로 반환한다.
func (t *JSONData) ToJSON() (*bytes.Buffer, error) {
	d, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(d)
	return b, nil
}

// Decode 함수는 JSONData 구조체로 디코딩한다.
func (t *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

// OtherJSONExamples 함수는 구조체나 다른 유용한 함수 이외의 다른 타입을 사용하는 방법을 보여준다.
func OtherJSONExamples() error {
	res := make(map[string]string)
	err := json.Unmarshal([]byte(`{"key": "value"}`), &res)
	if err != nil {
		return err
	}

	fmt.Println("We can unmarshal into a map instead of a struct:", res)

	b := bytes.NewReader([]byte(`{"key2": "value2"}`))
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(&res); err != nil {
		return err
	}

	fmt.Println("we can also use decoders/encoders to work with streams:", res)
	return nil
}
