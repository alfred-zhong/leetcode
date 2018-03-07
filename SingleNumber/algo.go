package algo

// https://leetcode.com/problems/single-number/description/
func singleNumber(nums []int) int {
	r := 0
	for _, v := range nums {
		r ^= v
	}
	return r
}
