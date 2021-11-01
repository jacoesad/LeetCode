package _75_distribute_candies

func distributeCandies(candyType []int) int {
	n := len(candyType)
	candyMap := make(map[int]bool, n/2)

	for _, candy := range candyType {
		candyMap[candy] = true
		if len(candyMap) >= n/2 {
			return n / 2
		}
	}

	return len(candyMap)
}
