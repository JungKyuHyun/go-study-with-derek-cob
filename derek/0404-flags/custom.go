package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CountTheWays 타입 선언
type CountTheWays []int

func (c *CountTheWays) String() string {
	result := ""
	// CountTheWays 배열만큼,
	for _, v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(v)
	}
	return result
}

// Split 한 문자열을 정수로 변환해 CountTheWays에 append
func (c *CountTheWays) Set(value string) error {
	values := strings.Split(value, ",")

	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*c = append(*c, i)
	}

	return nil
}
