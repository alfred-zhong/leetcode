package lswrc

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	max := 0
	sub := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		ii := findIndex(sub, s[i])
		sub = append(sub[ii+1:], s[i])

		if l := len(sub); l > max {
			max = l
		}
	}
	return max
}

func findIndex(bb []byte, b byte) int {
	for i := range bb {
		if bb[i] == b {
			return i
		}
	}
	return -1
}
