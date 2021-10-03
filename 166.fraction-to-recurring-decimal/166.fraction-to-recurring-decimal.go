func fractionToDecimal(numerator int, denominator int) string {
	intPart := strconv.Itoa(numerator / denominator)
	if numerator%denominator == 0 {
		return intPart
	}
	intPart += "."

	if (numerator < 0) != (denominator < 0) {
		if numerator < 0 {
			numerator = -numerator
		}
		if denominator < 0 {
			denominator = -denominator
		}
		intPart = "-" + strconv.Itoa(numerator/denominator) + "."
	}
	remainder := numerator % denominator
	decPart := ""
	remainderMap := make(map[int]int)

	for {
		_, exist := remainderMap[remainder]
		if exist {
			break
		}
		remainderMap[remainder] = len(decPart)

		curr := strconv.Itoa(remainder * 10 / denominator)
		remainder = remainder * 10 % denominator
		decPart += curr
		if remainder == 0 {
			break
		}
	}

	if remainder != 0 {
		decPart = decPart[:remainderMap[remainder]] + "(" + decPart[remainderMap[remainder]:] + ")"
	}

	return intPart + decPart
}
