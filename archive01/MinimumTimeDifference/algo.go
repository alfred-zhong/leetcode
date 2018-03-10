package algo

const period = 1440

func findMinDifference(timePoints []string) int {
	min := period / 2

	nums := make([]int, 0, len(timePoints))
	for _, t := range timePoints {
		i := timeToInt(t)
		insert := findInsert(nums, i)

		if insert >= 0 {
			newNums := make([]int, len(nums)+1)
			copy(newNums[:], nums[:insert])
			newNums[insert] = i
			copy(newNums[insert+1:], nums[insert:])
			nums = newNums

			if len(nums) > 1 {
				p, q := insert-1, insert+1
				if p < 0 {
					p = len(nums) - 1
				}
				if q >= len(nums) {
					q = 0
				}

				l := i - nums[p]
				if l < 0 {
					l += period
				}
				if l < min {
					min = l
				}

				r := nums[q] - i
				if r < 0 {
					r += period
				}
				if r < min {
					min = r
				}
			}
		} else {
			return 0
		}
	}

	return min
}

func timeToInt(time string) int {
	hour := int(time[0]-'0')*10 + int(time[1]-'0')
	minute := int(time[3]-'0')*10 + int(time[4]-'0')
	return hour*60 + minute
}

func findInsert(nums []int, i int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == i {
			return -1
		} else if nums[0] > i {
			return 0
		} else {
			return 1
		}
	}

	if i < nums[0] {
		return 0
	} else if i > nums[len(nums)-1] {
		return len(nums)
	}

	x, y := 0, len(nums)-1
	for y-x > 1 {
		medium := (x + y) / 2

		if nums[medium] == i {
			return -1
		} else if nums[medium] > i {
			y = medium
		} else {
			x = medium
		}
	}

	return y
}
