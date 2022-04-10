package templates

import (
	"fmt"
	"html/template"
	"os"
)

func HTMLDifferences() error {
	t := template.New("html")
	t, err := t.Parse("<h1>Hello! {{.Name}}</h1>\n")
	if err != nil {
		return err
	}

	// html/template 는 js의 주입(di) 동작을 자동 회피함.(변수 렌더링 위치에 따라 다름.)
	err = t.Execute(os.Stdout, map[string]string{"Name": "<script>alert('Can you see me?')</script>"})
	if err != nil {
		return err
	}

	// 수동으로 Escaper 실행.
	fmt.Println(template.JSEscaper(`example <example@example.com>`))
	fmt.Println(template.HTMLEscaper(`example <example@example.com>`))
	fmt.Println(template.URLQueryEscaper(`example <example@example.com>`))

	return nil
}
