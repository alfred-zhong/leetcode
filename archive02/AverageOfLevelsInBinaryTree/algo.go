// https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
package algo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	sums := make([]int, 0)
	cnts := make([]int, 0)

	iterateTree(root, 0, &sums, &cnts)

	r := make([]float64, len(sums))
	for i := range r {
		r[i] = float64(sums[i]) / float64(cnts[i])
	}

	return r
}

func iterateTree(n *TreeNode, level int, sums, cnts *[]int) {
	if n == nil {
		return
	}

	if len(*sums) <= level {
		*sums = append(*sums, 0)
		*cnts = append(*cnts, 0)
	}

	(*sums)[level] += n.Val
	(*cnts)[level]++

	iterateTree(n.Left, level+1, sums, cnts)
	iterateTree(n.Right, level+1, sums, cnts)
}
