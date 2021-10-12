package _9_divide_two_integers

import "math"

func divide(dividend int, divisor int) int {
	const INT_MAX = math.MaxInt32
	const INT_MIN = -INT_MAX - 1

	if dividend == INT_MIN {
		if divisor == 1 {
			return INT_MIN
		}
		if divisor == -1 {
			return INT_MAX
		}
	}

	if divisor == INT_MIN && dividend == INT_MIN {
		return 1
	}

	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	if divisor < dividend {
		return 0
	}

	candidates := []int{divisor}

	for candidates[len(candidates)-1] >= dividend-candidates[len(candidates)-1] {
		candidates = append(candidates, candidates[len(candidates)-1]+candidates[len(candidates)-1])
	}

	res := 0
	for i := len(candidates) - 1; i >= 0; i-- {
		if candidates[i] >= dividend {
			res += (1 << i)
			dividend -= candidates[i]
		}
	}

	if rev {
		res = -res
	}

	return res
}
