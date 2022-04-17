package collections

import "strings"

// 모든 데이터에 대해 ToLower
func LowerCaseData(w WorkWith) WorkWith {
	w.Data = strings.ToLower(w.Data)
	return w
}

// 버전 업
func IncrementVersion(w WorkWith) WorkWith {
	w.Version++
	return w
}

// old 버전인지 판별
func OldVersion(v int) func(w WorkWith) bool {
	return func(w WorkWith) bool {
		return w.Version >= v
	}
}