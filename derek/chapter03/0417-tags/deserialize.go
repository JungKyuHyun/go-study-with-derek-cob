package tags

import (
	"errors"
	"reflect"
	"strings"
)

// 직렬화된 문자열을 구조체로 변환한다.
func DeSerializeStructStrings(s string, res interface{}) error {
	r := reflect.TypeOf(res)

	// 포인터를 사용하기 때문에 포인터가 전달 됐는지 체크.
	if r.Kind() != reflect.Ptr {
		return errors.New("res must be a pointer")
	}

	// r = element type
	r = r.Elem()
	value := reflect.ValueOf(res).Elem()

	// 직렬화 데이터 split 
	vals := strings.Split(s, ";")
	valMap := make(map[string]string)
	// map 에 key value 나눠서 넣어줌.
	for _, v := range vals {
		keyval := strings.Split(v, ":")
		if len(keyval) != 2 {
			continue
		}
		valMap[keyval[0]] = keyval[1]
	}

	// r의 모든 필드 순회하며 재검토
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		// check serialize 필드
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			if serialize == "-" {
				continue
			}
			// serialize 필드면 "-" 빼고 스트링으로 바꿔줌.
			if val, ok := valMap[serialize]; ok {
				value.Field(i).SetString(val)
			}
		} else if val, ok := valMap[field.Name]; ok {
			value.Field(i).SetString(val)
		}
	}
	return nil
}