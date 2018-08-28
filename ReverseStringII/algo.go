// https://leetcode.com/problems/reverse-string-ii/
package algo

func reverseStr(s string, k int) string {
	r := []rune(s)
	for i := 0; i < len(r); i += k * 2 {
		if len(r[i:]) >= k {
			reverse(r[i : i+k])
		} else {
			reverse(r[i:])
		}
	}

	return string(r)
}

func reverse(r []rune) {
	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
}
