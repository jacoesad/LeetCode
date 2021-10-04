func licenseKeyFormatting(s string, k int) string {
	var builder strings.Builder
	var res strings.Builder
	n := len(s)
	m := 0
	builder.Grow(n)
	for i := 0; i < n; i++ {
		if s[i] == '-' {
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			builder.WriteByte(s[i] - 32)
		} else {
			builder.WriteByte(s[i])
		}
		m++
	}

	if m%k != 0 {
		res.Grow(m + m/k + 1)
		res.WriteString(builder.String()[0 : m%k])
		for i := m % k; i < m; i += k {
			res.WriteString("-")
			res.WriteString(builder.String()[i : i+k])
		}
	} else {
		res.Grow(m + m/k)
		for i := 0; i < m; i += k {
			if i != 0 {
				res.WriteString("-")
			}
			res.WriteString(builder.String()[i : i+k])
		}
	}

	return res.String()
}
