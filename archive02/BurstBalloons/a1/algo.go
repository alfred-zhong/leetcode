package algo

func maxCoins(nums []int) int {
	nn := make([]int, len(nums)+2)
	nn[0], nn[len(nn)-1] = 1, 1
	copy(nn[1:], nums)

	memo := make([][]int, len(nn))
	for i := range memo {
		memo[i] = make([]int, len(nn))
	}
	return burst(0, len(nn)-1, memo, nn)
}

func burst(left, right int, memo [][]int, nums []int) int {
	if right-left == 1 {
		return 0
	}

	if memo[left][right] > 0 {
		return memo[left][right]
	}

	max := 0
	for i := left + 1; i < right; i++ {
		t := nums[left]*nums[i]*nums[right] +
			burst(left, i, memo, nums) + burst(i, right, memo, nums)
		if t > max {
			max = t
		}
	}
	memo[left][right] = max
	return max
}
