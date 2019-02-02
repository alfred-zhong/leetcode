// https://leetcode.com/problems/contains-duplicate-iii/description/
package algo

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := len(nums)

	for x := 1; x <= k; x++ {
		for i := 0; i+x < l; i++ {
			c := nums[i] - nums[i+x]
			if c >= 0 && c <= t {
				return true
			} else if c < 0 && -c <= t {
				return true
			}
		}
	}

	return false
}
