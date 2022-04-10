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

// JSONData to  bytes.Buffer
func (t *JSONData) ToJSON() (*bytes.Buffer, error) {
	d, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)

	return b, nil
}

// Decode to JSONData
func (t *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

// 다른 데이터를 json으로 만드는 방법
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
