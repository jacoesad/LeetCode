package _73_integer_to_english_words

import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	var b strings.Builder

	if num/1e9 != 0 {
		b.WriteString(numberToWords(num / 1e9))
		b.WriteString(" Billion")
		num %= 1e9
	}

	if num == 0 {
		return b.String()
	}

	if b.String() != "" {
		b.WriteString(" ")
	}

	b.WriteString(numberToWordsLessThanBinllion(num))
	return b.String()
}

func numberToWordsLessThanBinllion(num int) string {
	var b strings.Builder

	if num/1e6 != 0 {
		b.WriteString(numberToWordsLessThanHundred(num/1e6) + " Million ")
		num = num % 1e6
	}

	if num/1e3 != 0 {
		b.WriteString(numberToWordsLessThanHundred(num/1e3) + " Thousand ")
		num = num % 1e3
	}

	if num != 0 {
		b.WriteString(numberToWordsLessThanHundred(num))
	}

	return strings.TrimSpace(b.String())
}

func numberToWordsLessThanHundred(num int) string {
	singles := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	doubles := []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	oneDoubles := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

	var b strings.Builder

	if num/100 != 0 {
		b.WriteString(singles[num/100] + " " + "Hundred" + " ")
		num %= 100
	}

	if num == 0 {
	} else if num < 10 {
		b.WriteString(singles[num])
	} else if num < 20 {
		b.WriteString(oneDoubles[num-10])
	} else {
		if num%10 == 0 {
			b.WriteString(doubles[num/10])
		} else {
			b.WriteString(doubles[num/10] + " " + singles[num%10])
		}
	}

	return strings.TrimSpace(b.String())
}
