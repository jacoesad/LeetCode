func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	n := len(s) - 10

	sMap := make(map[string]bool)
	res := []string{}

	for i := 0; i <= n; i++ {
		if _, ok := sMap[s[i:i+10]]; !ok {
			sMap[s[i:i+10]] = false
		} else {
			sMap[s[i:i+10]] = true
		}
	}

	for k, v := range sMap {
		if v == true {
			res = append(res, k)
		}
	}

	return res
}
