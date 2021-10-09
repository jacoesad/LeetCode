package _52_data_stream_as_disjoint_intervals

type SummaryRanges struct {
	NumList []int
}

func Constructor() SummaryRanges {
	return *new(SummaryRanges)
}

func (this *SummaryRanges) AddNum(val int) {
	n := len(this.NumList)
	if n == 0 {
		this.NumList = append(this.NumList, val)
		return
	}

	for i, num := range this.NumList {
		if val == num {
			return
		}

		if val < num {
			this.NumList = append(this.NumList[:i], append([]int{val}, this.NumList[i:]...)...)
			break
		}
		if i == n-1 {
			this.NumList = append(this.NumList, val)
		}
	}

	return
}

func (this *SummaryRanges) GetIntervals() [][]int {
	var res [][]int
	var tmp []int

	for _, num := range this.NumList {
		n := len(tmp)
		if n == 0 {
			tmp = append(tmp, num)
			continue
		}
		if num == tmp[n-1]+1 {
			tmp = append(tmp, num)
		} else {
			if n == 1 {
				tmp = append(tmp, tmp[0])
				res = append(res, tmp)
			} else {
				tmp = []int{tmp[0], tmp[n-1]}
				res = append(res, tmp)
			}
			tmp = []int{num}
		}
	}

	if tmp != nil {
		if len(tmp) == 1 {
			tmp = append(tmp, tmp[0])
		} else {
			tmp = []int{tmp[0], tmp[len(tmp)-1]}
		}
		res = append(res, tmp)
	}

	return res
}
