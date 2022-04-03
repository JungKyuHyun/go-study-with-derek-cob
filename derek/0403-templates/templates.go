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
        This is a way
        to insert a multi-line comment
    */}}
`

const secondTemplate = `
    {{ define "block_example" }}
        {{.OtherVariable}}
    {{end}}
`

func RunTemplate() error {
	// 익명 struct 선언 , 값 선언.
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
	// template에 함수를 등록.
	funcmap := template.FuncMap{
		"split": strings.Split,
	}
	// t = example 이라는 이름의 템플릿.
	t := template.New("example")
	// t의 함수를 funcmap 이라고 선언. (흡사 클래스 만드는 느낌)
	t = t.Funcs(funcmap)

	// Parse text as a template body for t
	t, err := t.Parse(sampleTemplate)
	if err != nil {
		return err
	}

	// 첫번째 템플릿을 복제.
	t2, err := t.Clone()
	if err != nil {
		return err
	}
	// 두번째 템플릿을 파싱.
	t2, err = t2.Parse(secondTemplate)
	if err != nil {
		return err
	}

	// 표준 출력에 템플릿을 출력하고 data를 씀.
	err = t2.Execute(os.Stdout, &data)
	if err != nil {
		return err
	}

	return nil
}
