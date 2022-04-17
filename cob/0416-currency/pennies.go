package currency

import "strconv"

// ConvertPenniesToDollarString함수는 int64로 페니의 야을 입력받아 달라의 문자열의 표현을 반환한다
func ConvertPenniesToDollarString(amount int64) string {
	// 10진수로 변환
	result := strconv.FormatInt(amount, 10)

	// 움수인 경우 나중에 설정
	negative := false
	if result[0] == '-' {
		result = result[1:]
		negative = true
	}

	// 값이 100보다 작으면 왼쪽에 0을 하나 추가한다.
	for len(result) < 3 {
		result = "0" + result
	}
	length := len(result)

	// 10진수로 덧셈한다
	result = result[0:length-2] + "." + result[length-2:]

	// 음수인 경우 "-"를 붙여준다
	if negative {
		result = "-" + result
	}
	return result
}
