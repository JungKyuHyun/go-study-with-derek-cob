package confformat

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// ToTOML함수는 TOMLData 구조체를 TOML 포맷의 bytes.Buffer로 반환한다.
func (t *TOMLData) ToTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)

	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return b, nil
}

// Decode 함수는 TOMLData 구조체로 디코딩한다.
func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}
