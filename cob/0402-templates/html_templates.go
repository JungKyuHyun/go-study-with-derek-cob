package templates

import (
	"fmt"
	"os"
	"text/template"
)

// HTMLDifferences 함수는 html/template과 text/template간의 차이점을 강조한다.
func HTMLDifferences() error {
	t := template.New("html")
	t, err := t.Parse("<h1>Hello! {{.Name}}</h1>n")
	if err != nil {
		return err
	}

	// html/template은 자바스크립트 주입과 같은 안전하지 않은 동작을
	// 자동 회피(escape) 처리한다. 이는 상황에 따라 인식되며, 변수가 렌더링되는 위치에 따라 다르게 동작한다.
	err = t.Execute(os.Stdout, map[string]string{"Name": "<script>alert('Can you see me?')</script>"})
	if err != nil {
		return err
	}

	// 수동으로 escaper를 호출할 수도 있다.
	fmt.Println(template.JSEscaper(`example <example@example.com>`))
	fmt.Println(template.HTMLEscaper(`example <example@example.com>`))
	fmt.Println(template.URLQueryEscaper(`example <example@example.com>`))

	return nil
}
