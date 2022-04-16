package currency

import (
	"errors"
	"strconv"
	"strings"
)

// ConvertStringDollarsToPennies함수는 문자열로 달러의 양을 입력받아 int64를 반환한다.
func ConvertStringDollarsToPennies(amount string) (int64, error) {
	// amount 인자가 float로 변환 가능한지 확인
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	// . 문자를 기준으로 값을 나눈다(분할)
	groups := strings.Split(amount, ".")

	// . 문자가 없는 경우에는 다음 코드에서 값을 result에 저장한다
	result := groups[0]

	// 기본 문자열
	r := ""

	// "." 다음의 데이터를 처리
	if len(groups) == 2 {
		if len(groups[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = groups[1]
	}
	// 0으로 채움
	//  . 문자가 없는 경우 두개의 0으로 채운다
	for len(r) < 2 {
		r += "0"
	}
	result += r

	return strconv.ParseInt(result, 10, 64)
}
