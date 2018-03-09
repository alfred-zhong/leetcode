package algo

// https://leetcode.com/problems/reach-a-number/description/
func reachNumber(target int) int {
	if target == 0 {
		return 0
	}

	t := target
	if t < 0 {
		t = -t
	}

	k := 0
	sum := 0
	for {
		k++
		sum += k

		if sum >= t && (sum-t)%2 == 0 {
			break
		}
	}
	return k
}
