package algo

// https://leetcode.com/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sum[i] = nums[i]
		} else {
			sum[i] = sum[i-1] + nums[i]
		}
	}

	max := sum[0]
	min := 0
	for j := 1; j < len(nums); j++ {
		if sum[j-1] < min {
			min = sum[j-1]
		}

		if t := sum[j] - min; t > max {
			max = t
		}
	}

	return max
}
