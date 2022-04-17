package collections

type WorkWith struct {
	Data    string
	Version int
}

// Filter를 직접 구현. 
func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {
	// WorkWith 배열 길이 0 으로 초기화
	result := make([]WorkWith, 0)
	for _, w := range ws {
		if f(w) {
			// WorkWith 객체가 bool 이라면 result에 append
			result = append(result, w)
		}
	}
	return result
}

// Map 구현
func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {
	result := make([]WorkWith, len(ws))

	for pos, w := range ws {
		newW := f(w)
		result[pos] = newW
		// map 으로 매핑.
	}
	return result
}