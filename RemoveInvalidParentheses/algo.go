package algo

// Just from https://leetcode.com/problems/remove-invalid-parentheses/discuss/75027/Easy-Short-Concise-and-Fast-Java-DFS-3-ms-solution.
// So brilliant. How can I think of such an idea.
func removeInvalidParentheses(s string) []string {
	ans := make([]string, 0)
	remove(s, &ans, 0, 0, []byte{'(', ')'})
	return ans
}

func remove(s string, ans *[]string, lastI, lastJ int, par []byte) {
	for stack, i := 0, lastI; i < len(s); i++ {
		if s[i] == par[0] {
			stack++
		}
		if s[i] == par[1] {
			stack--
		}
		if stack >= 0 {
			continue
		}

		for j := lastJ; j <= i; j++ {
			if s[j] == par[1] && (j == lastJ || s[j-1] != par[1]) {
				remove(s[0:j]+s[j+1:len(s)], ans, i, j, par)
			}
		}
		return
	}

	reversed := reverseString(s)
	if par[0] == '(' {
		remove(reversed, ans, 0, 0, []byte{')', '('})
	} else {
		*ans = append(*ans, reversed)
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
