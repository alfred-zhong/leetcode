package a1

func longestPalindrome(s string) string {
	r := ""
	l := len(s)
	for i := 0; i < l-len(r); i++ {
		for j := l - 1; j > len(r)+i-1; j-- {
			if isPalindrome(s[i : j+1]) {
				r = s[i : j+1]
			}
		}
	}
	return r
}

func isPalindrome(s string) bool {
	if s == "" {
		return false
	}

	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
