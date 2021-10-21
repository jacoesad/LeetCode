package _53_minimum_moves_to_equal_array_elements

func minMoves(nums []int) int {
	min := nums[0]
	sum := 0
	for _, num := range nums {
		sum += num
		if num < min {
			min = num
		}
	}

	return sum - min*len(nums)
}
