func toHex(num int) string {
	if num > 0 {
		return posToHex(num)
	}

	if num < 0 {
		return nagToHex(num)
	}

	return "0"
}

func posToHex(num int) (s string) {
	hexMap := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}

	for num != 0 {
		tmp := num % 16
		num = num / 16
		s = hexMap[tmp] + s
	}

	return
}

func nagToHex(num int) (s string) {
	binNum := []byte("00000000000000000000000000000000")
	absNum := -num
	hexNum := ""
	hexMap := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}

	for i := len(binNum) - 1; i >= 0; i-- {
		if absNum%2 == 0 {
			binNum[i] = '1'
		}
		absNum /= 2
	}

	for i := len(binNum) - 1; i > 0; i-- {
		if binNum[i] == '1' {
			binNum[i] = '0'
		} else {
			binNum[i] = '1'
			break
		}
	}
	binNum[0] = '1'

	for i := 0; i < 32; i = i + 4 {
		num := 0
		for j := 0; j < 4; j++ {
			if binNum[i+j] == '1' {
				num += pow(2, 3-j)
			}

		}
		hexNum += hexMap[num]
	}

	return hexNum
}

func pow(x int, n int) int {
	if n == 0 {
		return 1
	}

	return x * pow(x, n-1)
}
