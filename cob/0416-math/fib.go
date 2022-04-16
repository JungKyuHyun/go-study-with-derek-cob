package math

import "math/big"

// 피보나치 수열의 저장을 위한 글로벌 변수
var memoize map[int]*big.Int

func init() {
	// map 초기화
	memoize = make(map[int]*big.Int)
}

// Fib함수는 피보나치 수열의 n번째 수를 출력한다.
// 0보다 작은 경우에는 1를 반환
// 재귀적으로 계산하며, int64는 빨리 오버플로되기 때문에
// big.Int를 사용한다.
func Fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}

	// 기본 케이스
	if n < 2 {
		memoize[n] = big.NewInt(1)
	}

	// 이전에 저장했는지 확인하고,
	// 이미 저장한 경우, 계산하지 않고 저장된 값을 반환
	if val, ok := memoize[n]; ok {
		return val
	}

	// 맵을 초기화한 다음, 이전 두 개의 피보나치 값을 더한다.
	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], Fib(n-1))
	memoize[n].Add(memoize[n], Fib(n-2))

	return memoize[n]
}
