package _76_number_complement

func findComplement(num int) int {
	highBit := 1

	for 1<<highBit <= num {
		highBit++
	}

	return num ^ (1<<(highBit) - 1)
}
