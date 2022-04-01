package templates

import (
	"os"
	"strings"
	"text/template"
)

const sampleTemplate = `
    This template demonstrates printing a {{ .Variable | printf "%#v" }}.
    
    {{if .Condition}}
    If condition is set, we'll print this
    {{else}}
    Otherwise, we'll print this instead
    {{end}}
    Next we'll iterate over an array of strings:
    {{range $index, $item := .Items}}
        {{$index}}: {{$item}}
    {{end}}
    We can also easily import other functions like strings.Split
    then immediately used the array created as a result:
    {{ range $index, $item := split .Words ","}}
        {{$index}}: {{$item}}
    {{end}}
    Blocks are a way to embed templates into one another
    {{ block "block_example" .}}
        No Block defined!
    {{end}}
    {{/*
        여러 줄 주석을 추가하는 방법
    */}}
`

const secondTemplate = `
    {{ define "block_example" }}
        {{.OtherVariable}}
    {{end}}
`

// RunTemplate 함수는 템플릿을 초기화하고 다양한 템플릿 도움 함수의 예시를 보여준다.
func RunTemplate() error {
	data := struct {
		Condition     bool
		Variable      string
		Items         []string
		Words         string
		OtherVariable string
	}{
		Condition:     true,
		Variable:      "variable",
		Items:         []string{"item1", "item2", "item3"},
		Words:         "another_item1,another_item2,another_item3",
		OtherVariable: "I'm defined in a second template!",
	}
	funcmap := template.FuncMap{
		"split": strings.Split,
	}

	t := template.New("example")
	t = t.Funcs(funcmap)
	// 오류에 대응하기 위해 Must를 대신 사용할 수 있따.
	// template.Must(t.Parse(sampleTemplate))
	t, err := t.Parse(sampleTemplate) 
	if err != nil {
		return err
	}

	// 블록의 예시를 보여주기 위해 첫 번째 템플릿을 복제한 다음
	// 두 번째 템플릿을 파싱해 또 다른 템플릿을 생성한다.
	t2, err := t.Clone()
	if err != nil {
		return err
	}
	t2, err = t2.Parse(secondTemplate)
	if err != nil {
		return err
	}
	// 표준 출력에 템플릿을 출력하고 데이터로 채운다.
	err = t2.Execute(os.Stdout, &data)
	if err != nil {
		return err
	}
	return nil
}