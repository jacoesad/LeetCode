type MyCalendarTwo struct {
	// 记录所有预订的时间段
	bookings [][2]int
	// 记录所有双重预订的时间段
	overlaps [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		bookings: make([][2]int, 0),
		overlaps: make([][2]int, 0),
	}
}

func (cal *MyCalendarTwo) Book(start int, end int) bool {
	// 首先检查是否会与任何双重预订重叠
	for _, overlap := range cal.overlaps {
		if start < overlap[1] && end > overlap[0] {
			return false // 会导致三重预订
		}
	}

	// 检查与现有预订的重叠，并记录重叠部分
	for _, booking := range cal.bookings {
		if start < booking[1] && end > booking[0] {
			// 记录重叠区间
			overlapStart := max(start, booking[0])
			overlapEnd := min(end, booking[1])
			cal.overlaps = append(cal.overlaps, [2]int{overlapStart, overlapEnd})
		}
	}

	// 添加新的预订
	cal.bookings = append(cal.bookings, [2]int{start, end})
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
// @lc code=end
