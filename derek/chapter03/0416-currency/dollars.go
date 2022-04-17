package currency

import (
	"errors"
	"strconv"
	"strings"
)

func ConvertStringDollarsToPennies(amount string) (int64, error) {
	// amount를 float로 변환
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}
	// .으로 분리
	groups := strings.Split(amount, ".")

	result := groups[0]
	r := ""

	if len(groups) == 2 {
		if len(groups[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = groups[1]
	}

	for len(r) < 2 {
		r += "0"
	}

	result += r

	return strconv.ParseInt(result, 10, 64)
}