package confformat

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// TOMLData 포멧을 TOML bytes buffer로 반환함.
func (t *TOMLData) ToTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)

	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return b, nil
}

// Decode TOMLData
func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}
