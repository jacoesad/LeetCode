package _67_valid_perfect_square

func isPerfectSquare(num int) bool {
	for i := 0; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
