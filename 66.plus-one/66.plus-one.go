package _6_plus_one

func plusOne(digits []int) []int {
	n := len(digits) - 1
	addOne := true

	for addOne && n >= 0 {
		digits[n] += 1
		if digits[n] != 10 {
			addOne = false
			break
		}
		digits[n] = 0
		n--
	}

	if addOne {
		digits = append([]int{1}, digits...)
	}

	return digits
}
