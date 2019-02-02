// https://leetcode.com/problems/maximum-binary-tree/description/
package algo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max, index := findMaxIndex(nums)
	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}

func findMaxIndex(nums []int) (max, index int) {
	max = nums[0]
	index = 0
	for i, num := range nums {
		if num > max {
			index = i
			max = num
		}
	}
	return
}
