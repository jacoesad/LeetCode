func thirdMax(nums []int) int {
	res := make([]*int, 3)
	for _, num := range nums {
		num := num
		if res[0] == nil {
			res[0] = &num
		} else if res[1] == nil && num > *res[0] {
			res[1] = &num
		} else if res[1] == nil && num < *res[0] {
			res[0], res[1] = &num, res[0]
		} else if res[1] == nil && num == *res[0] {
			continue
		} else if res[2] == nil && num > *res[0] && num > *res[1] {
			res[2] = &num
		} else if res[2] == nil && num > *res[0] && num < *res[1] {
			res[1], res[2] = &num, res[1]
		} else if res[2] == nil && num < *res[0] && num < *res[1] {
			res[0], res[1], res[2] = &num, res[0], res[1]
		} else if res[2] == nil && (num == *res[0] || num == *res[1]) {
			continue
		} else if num > *res[1] && num > *res[2] {
			res[0], res[1], res[2] = res[1], res[2], &num
		} else if num > *res[1] && num < *res[2] {
			res[0], res[1] = res[1], &num
		} else if num > *res[0] && num < *res[1] {
			res[0] = &num
		}
	}

	if res[2] == nil && res[1] != nil {
		return *res[1]
	}

	return *res[0]
}
