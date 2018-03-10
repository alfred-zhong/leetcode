// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	r := make([][]int, 0)
	levelStore(root, 0, &r)

	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	return r
}

func levelStore(n *TreeNode, level int, m *[][]int) {
	if len(*m) <= level {
		*m = append(*m, make([]int, 0))
	}

	(*m)[level] = append((*m)[level], n.Val)
	if n.Left != nil {
		levelStore(n.Left, level+1, m)
	}
	if n.Right != nil {
		levelStore(n.Right, level+1, m)
	}
}
