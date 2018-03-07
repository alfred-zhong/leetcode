package a2

func longestPalindrome(s string) string {
	r := ""

	m := len(s) / 2
	for i := m; i >= len(r)/2; i-- {
		if ss := longestIndexCenter(i, i, s); len(ss) > len(r) {
			r = ss
		}
		if i+1 < len(s) {
			if ss := longestIndexCenter(i, i+1, s); len(ss) > len(r) {
				r = ss
			}
		}
	}
	for i := m + 1; i < len(s)-len(r)/2; i++ {
		if ss := longestIndexCenter(i, i, s); len(ss) > len(r) {
			r = ss
		}
		if i+1 < len(s) {
			if ss := longestIndexCenter(i, i+1, s); len(ss) > len(r) {
				r = ss
			}
		}
	}

	return r
}

func longestIndexCenter(x, y int, s string) string {
	for x >= 0 && y < len(s) {
		if s[x] != s[y] {
			return s[x+1 : y]
		}

		x--
		y++
	}
	return s[x+1 : y]
}
