// https://leetcode.com/problems/next-greater-element-iii/description/
package algo

import (
	"sort"
)

func nextGreaterElement(n int) int {
	digits := int2Digits(n)
	if len(digits) == 1 {
		return -1
	}

	for i := len(digits) - 2; i >= 0; i-- {
		if !isDecreased(digits[i:]) {
			adjust(digits[i:])
			r := digits2Int(digits)
			if r > 1 << 31 {
				return -1
			}

			return r
		}
	}

	return -1
}

func adjust(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	// find number greater than nums[0] and exchange
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[0] {
			nums[0], nums[i] = nums[i], nums[0]
			break
		}
	}

	sort.Ints(nums[1:])
}

func isDecreased(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}

func int2Digits(i int) []int {
	r := make([]int, 0, 19)
	for i > 0 {
		r = append(r, i%10)
		i /= 10
	}
	reverseInts(r)
	return r
}

func digits2Int(d []int) int {
	if d == nil || len(d) == 0 {
		return 0
	}

	r := 0
	for _, i := range d {
		r = r*10 + i
	}
	return r
}

func reverseInts(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
