package twoSum

func twoSum(nums []int, target int) []int {
	m := make(map[int][]int, len(nums))
	for i, v := range nums {
		if indexes, ok := m[v]; ok {
			indexes := append(indexes, i)
			m[v] = indexes
		} else {
			m[v] = []int{i}
		}
	}

	for k1, v1 := range m {
		k2 := target - k1

		if k1 == k2 {
			if len(v1) > 1 {
				return []int{v1[0], v1[1]}
			}
			continue
		}

		if v2, ok := m[k2]; ok {
			return []int{v1[0], v2[0]}
		}
	}

	return []int{}
}
