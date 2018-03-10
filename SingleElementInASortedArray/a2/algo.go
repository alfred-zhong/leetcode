package a2

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	i, j := 0, len(nums)-1

	for i != j {
		m := (i + j) / 2

		switch {
		case nums[m] != nums[m-1] && nums[m] != nums[m+1]:
			return nums[m]
		case nums[m] == nums[m-1]:
			if (m-i)%2 == 0 {
				j = m - 2
			} else {
				i = m + 1
			}
		case nums[m] == nums[m+1]:
			if (m-i)%2 == 0 {
				i = m + 2
			} else {
				j = m - 1
			}
		}
	}

	return nums[i]
}
