// https://leetcode.com/problems/intersection-of-two-arrays/description/
package algo

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, n := range nums1 {
		m[n] = true
	}

	r := make([]int, 0, min(len(nums1), len(nums2)))
	for _, n := range nums2 {
		if m[n] {
			r = append(r, n)
			delete(m, n)
		}
	}

	return r
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
