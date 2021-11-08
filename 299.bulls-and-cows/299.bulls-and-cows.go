package _99_bulls_and_cows

import "strconv"

func getHint(secret string, guess string) string {
	n := len(secret)
	secretMap := make([]int, 10)
	guessMap := make([]int, 10)
	bull := 0
	cow := 0
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			bull++
			continue
		}
		secretMap[int(secret[i]-48)]++
		guessMap[int(guess[i])-48]++
	}
	for i := 0; i <= 9; i++ {
		cow += min(secretMap[i], guessMap[i])
	}

	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
