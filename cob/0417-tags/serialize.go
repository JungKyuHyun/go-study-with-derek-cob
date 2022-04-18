package tags

import "reflect"

// SerializeStructStrings 함수는 구조체를 직렬화 포맷으로 변환하고, 문자열 타입에 대한 구조체 태크를 직렬화하는 방식을 사용한다.
func SerializeStructStrings(s interface{}) (string, error) {
	result := ""

	// 인터페이스를 특정 타입으로 반영한다.
	r := reflect.TypeOf(s)
	value := reflect.ValueOf(s)

	// 구조체에 대한 포인터가 전달된 경우 적절하게 처리한다.
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}

	// 모든 필드에 대해 루프로 처리
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		// 구조체 태그가 발견된 경우
		key := field.Name
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// "-"를 무시한다. 그렇지 않으면 전체 값이 직렬화 키가 된다.
			if serialize == "-" {
				continue
			}
			key = serialize
		}

		switch value.Field(i).Kind() {
		// 이 예제는 문자열만 지원
		case reflect.String:
			result += key + ":" + value.Field(i).String() + ";"
		// 이외 스킵
		default:
			continue
		}
	}
	return result, nil
}
