// https://leetcode.com/submissions/detail/164623311/
package algo

import (
	"sort"
)

func findKthLargest(nums []int, k int) int {
	r := make([]int, k)
	copy(r, nums)
	sort.Ints(r)

	for i := k; i < len(nums); i++ {
		insert(r, nums[i])
	}

	return r[0]
}

func insert(r []int, num int) {
	if len(r) == 0 {
		return
	}

	if num <= r[0] {
		return
	}

	for i := len(r) - 1; i >= 0; i-- {
		if num > r[i] {
			copy(r, r[1:i+1])
			r[i] = num
			break
		}
	}
}
