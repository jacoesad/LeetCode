func countSegments(s string) int {
	res := 0
	pre := false
	for _, b := range s {
		if b == ' ' {
			pre = false
		} else if pre == false {
			res++
			pre = true
		}
	}

	return res
}
