package algo

import (
	"strconv"
)

// https://leetcode.com/problems/construct-string-from-binary-tree/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	ss := strconv.Itoa(t.Val)
	if t.Left != nil && t.Right != nil {
		ss += "(" + tree2str(t.Left) + ")(" + tree2str(t.Right) + ")"
	} else if t.Left != nil && t.Right == nil {
		ss += "(" + tree2str(t.Left) + ")"
	} else if t.Left == nil && t.Right != nil {
		ss += "()(" + tree2str(t.Right) + ")"
	}

	return ss
}
