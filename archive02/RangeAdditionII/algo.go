// https://leetcode.com/problems/range-addition-ii/description/
package algo

func maxCount(m int, n int, ops [][]int) int {
	x, y := m, n
	for _, op := range ops {
		x, y = intersect(x, y, op[0], op[1])
	}

	return x * y
}

func intersect(m, n int, x, y int) (int, int) {
	if m <= 0 || n <= 0 {
		return 0, 0
	}

	if x <= 0 || y <= 0 {
		return m, n
	}

	return min(m, x), min(n, y)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
