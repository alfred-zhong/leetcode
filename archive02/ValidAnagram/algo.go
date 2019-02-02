// https://leetcode.com/problems/valid-anagram/description/
package algo

func isAnagram(s string, t string) bool {
	cnts := make([]int, 128)

	for _, ss := range s {
		cnts[ss]++
	}

	for _, tt := range t {
		cnts[tt]--
	}

	for _, cnt := range cnts {
		if cnt != 0 {
			return false
		}
	}

	return true
}
