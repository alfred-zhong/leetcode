package algo

// https://leetcode.com/problems/reach-a-number/description/
func reachNumber(target int) int {
	if target == 0 {
		return 0
	}

	if target < 0 {
		target = -target
	}

	k := 0
	sum := 0
	for {
		k++
		sum += k

		if sum >= target && (sum-target)%2 == 0 {
			break
		}
	}
	return k
}
