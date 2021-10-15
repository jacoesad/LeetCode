package 剑指_Offer_II_069__山峰数组的顶部

func peakIndexInMountainArray(arr []int) int {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return i - 1
		}
	}
	return 0
}
