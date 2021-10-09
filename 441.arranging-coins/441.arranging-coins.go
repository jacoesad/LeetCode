package _41_arranging_coins

func arrangeCoins(n int) int {
	for i := 1; ; i++ {
		if i == n {
			return i
		}
		if i > n {
			return i - 1
		}
		n -= i
	}
}
