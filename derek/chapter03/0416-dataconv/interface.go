package dataconv

import "fmt"

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
// 타입이 구현해야 하는 메서드의 원형.
func Interfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)

	var i interface{}
	i = "test"
	// Type Assertion : i는 string 타입에 속한다는 점을 확인(assert)하는 것
	if val, ok := i.(string); ok {
		fmt.Println("val is", val)
	}

	if _, ok := i.(int); !ok {
		fmt.Println("uh oh! glad we handled this")
	}
}