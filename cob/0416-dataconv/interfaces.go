package dataconv

import "fmt"

// CheckType함수는 인터페이스 타입을 기반으로 출력한다.
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("It's a string!")
	case int:
		fmt.Println("It's an int!")
	default:
		fmt.Println("not sure what it is...")
	}
}

// Interfaces함수는 익명의 인터페이스를 다른 타입으로 형 변환하는 방법을 보여준다.
func Interfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)

	var i interface{}
	i = "test"

	// 인터페이스를 수동으로 확인
	if val, ok := i.(string); ok {
		fmt.Println("val is", val)
	}

	// 실패 케이스
	if _, ok := i.(int); !ok {
		fmt.Println("un oh! glad we handled this")
	}
}
