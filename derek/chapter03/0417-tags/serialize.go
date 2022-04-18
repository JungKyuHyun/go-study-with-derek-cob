package tags

// 리플렉션은 실행 시점(Runtime, 런타임)에 인터페이스나 구조체 등의 타입 정보를 얻어내거나 결정하는 기능
import "reflect"

func SerializeStructStrings(s interface{}) (string, error) {
	result := ""

	// 인터페이스 s를 특정 타입으로 반영한다
	r := reflect.TypeOf(s)
	// 인터페이스 s의 값을 value에 넣음.
	value := reflect.ValueOf(s)

	// r의 타입이 reflect 포인터라면,
	if r.Kind() == reflect.Ptr {
		// r type's element type 을 리턴.
		r = r.Elem()
		value = value.Elem()
	}

	// r의 모든 필드에 대해 루프를 돈다
	for i := 0; i < r.NumField(); i++ {
		// i'th field 
		field := r.Field(i)

		key := field.Name
		// return "serialize" associated with key in the tag string.
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// "-" 는 무시한다 왜냐하면 전체값이 직렬화 키가 되는것을 방지하기 위햄
			if serialize == "-" {
				continue
			}

			key = serialize
		}

		switch value.Field(i).Kind() {
		// value의 i번째 필드의 type을 value
		case reflect.String:
			// 스트링일 경우
			result += key + ":" + value.Field(i).String() + ";"
		default:
			continue
		}
	}
	return result, nil
}