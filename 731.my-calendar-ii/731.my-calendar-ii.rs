struct MyCalendarTwo {
    // 记录所有预订的时间段
    bookings: Vec<(i32, i32)>,
    // 记录所有双重预订的时间段
    overlaps: Vec<(i32, i32)>,
}

impl MyCalendarTwo {
    fn new() -> Self {
        MyCalendarTwo {
            bookings: Vec::new(),
            overlaps: Vec::new(),
        }
    }
    
    fn book(&mut self, start: i32, end: i32) -> bool {
        // 首先检查是否会与任何双重预订重叠
        for &(overlap_start, overlap_end) in &self.overlaps {
            if start < overlap_end && end > overlap_start {
                return false; // 会导致三重预订
            }
        }
        
        // 检查与现有预订的重叠，并记录重叠部分
        for &(booking_start, booking_end) in &self.bookings {
            if start < booking_end && end > booking_start {
                // 记录重叠区间
                let overlap_start = start.max(booking_start);
                let overlap_end = end.min(booking_end);
                self.overlaps.push((overlap_start, overlap_end));
            }
        }
        
        // 添加新的预订
        self.bookings.push((start, end));
        true
    }
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * let obj = MyCalendarTwo::new();
 * let ret_1: bool = obj.book(start, end);
 */
