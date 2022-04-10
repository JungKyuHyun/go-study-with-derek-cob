package confformat

import (
	"bytes"

	"github.com/go-yaml/yaml"
)

type YAMLData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

// ToYAML 함수는 YAMLData 구조체를 YAML 포맷의 bytes.Buffer로 반환한다.
func (t *YAMLData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(d)
	return b, nil
}

// Decode 함수는 YAMLData 구조체로 디코딩한다.
func (t *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}
