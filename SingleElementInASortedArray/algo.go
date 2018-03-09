// https://leetcode.com/problems/single-element-in-a-sorted-array/description/
package algo

func singleNonDuplicate(nums []int) int {
	r := 0
	for i := 0; i < len(nums); {
		r ^= nums[i]
		i++

		if i >= len(nums) {
			return nums[i-1]
		}

		r ^= nums[i]
		if r != 0 {
			return nums[i-1]
		}
		i++
	}
	return r
}
