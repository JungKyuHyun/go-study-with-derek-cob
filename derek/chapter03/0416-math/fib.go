package math

import "math/big"

var memoize map[int]*big.Int
// init() 함수는 패키지가 로드될 때 가장 먼저 호출되는 함수로, 패키지의 초기화 로직이 필요할 때 선택적으로 사용하면 된다.
func init() {
	memoize = make(map[int]*big.Int)
}

func Fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}

	if n < 2 {
		memoize[n] = big.NewInt(1)
	}
	// 이미 저장된 값이 있는 경우 리턴.
	if val, ok := memoize[n]; ok {
		return val
	}

	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], Fib(n-1))
	memoize[n].Add(memoize[n], Fib(n-2))

	return memoize[n]
}