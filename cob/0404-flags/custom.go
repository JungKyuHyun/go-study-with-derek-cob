package flags

import (
	"fmt"
	"strconv"
	"strings"
)

// CountTheWays는 플래그를 읽어들일 커스텀 타입
type CountTheWays []int

func (c *CountTheWays) String() string {
	result := ""
	for _, v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(v)
	}
	return result
}

// Set 함수는 flag 패키지에서 사용할 될 예정
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
