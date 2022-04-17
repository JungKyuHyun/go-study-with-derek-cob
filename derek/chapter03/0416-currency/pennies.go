package currency

import (
	"strconv"
)

func ConvertPenniesToDollarString(amount int64) string {
	result := strconv.FormatInt(amount, 10)
	// 음수 체크
	negative := false
	if result[0] == '-' {
		result = result[1:]
		negative = true
	}

	// 100보다 작으면 1달러가 아니니까 왼쪽에 0 추가.
	for len(result) < 3 {
		result = "0" + result
	}
	length := len(result)

	// 배열은 파이썬 처럼 [start:end] 방식 가능
	result = result[0:length-2] + "." + result[length-2:]

	if negative {
		result = "-" + result
	}

	return result
}