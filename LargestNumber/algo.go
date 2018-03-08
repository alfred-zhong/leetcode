package algo

import (
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.com/submissions/detail/144043940/
func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i := range nums {
		ss[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(ss, func(i, j int) bool {
		x, y := ss[i]+ss[j], ss[j]+ss[i]
		return x < y
	})

	result := ""
	for i := len(ss) - 1; i >= 0; i-- {
		result += ss[i]
	}
	result = trimZero(result)

	return result
}

func trimZero(s string) string {
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return "0"
	}
	return s
}
