package templates

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

// CreateTemplate 함수는 데이터를 포함하는 템플릿 파일을 생성한다.
func CreateTemplate(path string, data string) error {
	return ioutil.WriteFile(path, []byte(data), os.FileMode(0755))
}

// InitTemplates 함수는 디렉터리로부터 템플릿을 설정한다.
func InitTemplates() error {
	tempdir, err := ioutil.TempDir("", "temp")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempdir)

	err = CreateTemplate(filepath.Join(tempdir, "t1.tmpl"), `
        Template 1! {{ .Var1 }}
        {{ block "template2" .}} {{end}}
        {{ block "template3" .}} {{end}}
    `)
	if err != nil {
		return err
	}

	err = CreateTemplate(filepath.Join(tempdir, "t2.tmpl"), `
        {{ define "template2"}}Template 2! {{ .Var2 }}{{end}}
    `)
	if err != nil {
		return err
	}

	err = CreateTemplate(filepath.Join(tempdir, "t3.tmpl"), `
        {{ define "template3"}}Template 3! {{ .Var3 }}{{end}}
    `)
	if err != nil {
		return err
	}

	pattern := filepath.Join(tempdir, "*.tmpl")
	// ParseGlob 함수는 패턴(조건)과 일치하는 모든 파일을 모아 하나의 템플릿으로 결합시킨다.
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return err
	}
	// Execute 함수는 구조체 대신 맵(map)타입에도 사용할 수 있다.
	tmpl.Execute(os.Stdout, map[string]string{
		"Var1": "Var1!!",
		"Var2": "Var2!!",
		"Var3": "Var3!!",
	})
	return nil
}
