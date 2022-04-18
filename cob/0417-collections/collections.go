package collections

type WorkWith struct {
	Data    string
	Version int
}

// Filter는 함수형 필터다
func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {
	result := make([]WorkWith, 0)
	for _, w := range ws {
		if f(w) {
			result = append(result, w)
		}
	}
	return result
}

// Map함수는 함수형 맵이다.
func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {
	result := make([]WorkWith, len(ws))

	for pos, w := range ws {
		newW := f(w)
		result[pos] = newW
	}
	return result
}
