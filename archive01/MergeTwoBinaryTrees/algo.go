// https://leetcode.com/problems/merge-two-binary-trees/description/
package algo

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
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 != nil {
		t3 := new(TreeNode)
		t3.Val = t1.Val + t2.Val
		t3.Left = mergeTrees(t1.Left, t2.Left)
		t3.Right = mergeTrees(t1.Right, t2.Right)
		return t3
	} else if t1 != nil && t2 == nil {
		t3 := new(TreeNode)
		t3.Val = t1.Val
		t3.Left = mergeTrees(t1.Left, nil)
		t3.Right = mergeTrees(t1.Right, nil)
		return t3
	} else if t1 == nil && t2 != nil {
		t3 := new(TreeNode)
		t3.Val = t2.Val
		t3.Left = mergeTrees(nil, t2.Left)
		t3.Right = mergeTrees(nil, t2.Right)
		return t3
	}

	return nil
}
