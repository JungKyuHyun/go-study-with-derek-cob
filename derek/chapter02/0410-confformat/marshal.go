package confformat

import "fmt"

// 구조체로 된 데이터를 다양한 데이터 포멧으로 변환해줌
func MarshalAll() error {
	t := TOMLData{
		Name: "Name1",
		Age:  20,
	}

	j := JSONData{
		Name: "Name2",
		Age:  30,
	}

	y := YAMLData{
		Name: "Name3",
		Age:  40,
	}

	tomlRes, err := t.ToTOML()
	if err != nil {
		return err
	}

	fmt.Println("TOML Marshal =", tomlRes.String())

	jsonRes, err := j.ToJSON()
	if err != nil {
		return err
	}

	fmt.Println("JSON Marshal=", jsonRes.String())

	yamlRes, err := y.ToYAML()
	if err != nil {
		return err
	}

	fmt.Println("YAML Marshal =", yamlRes.String())
	return nil
}
