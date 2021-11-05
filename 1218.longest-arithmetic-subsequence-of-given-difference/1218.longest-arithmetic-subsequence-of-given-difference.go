package _218_longest_arithmetic_subsequence_of_given_difference

func longestSubsequence(arr []int, difference int) int {
	res := 0
	dp := map[int]int{}

	for _, num := range arr {
		dp[num] = dp[num-difference] + 1
		if dp[num] > res {
			res = dp[num]
		}
	}

	return res
}
