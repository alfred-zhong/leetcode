// https://leetcode.com/problems/burst-balloons/description/
package algo

func maxCoins(nums []int) int {
	nn := make([]int, len(nums)+2)
	n := len(nn)
	nn[0], nn[n-1] = 1, 1
	copy(nn[1:], nums)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	for k := 2; k < n; k++ {
		for left := 0; left < n-k; left++ {
			right := left + k
			for i := left + 1; i < right; i++ {
				t := nn[left]*nn[i]*nn[right] +
					memo[left][i] + memo[i][right]
				if t > memo[left][right] {
					memo[left][right] = t
				}
			}
		}
	}

	return memo[0][n-1]
}
