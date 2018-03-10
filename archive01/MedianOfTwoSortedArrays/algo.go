package algo

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1, p2 := 0, 0
	sumLength := len(nums1) + len(nums2)
	if sumLength%2 == 0 {
		p1, p2 = sumLength/2-1, sumLength/2
	} else {
		p1 = sumLength / 2
		p2 = p1
	}

	sum := 0
Loop:
	for i := 0; i < sumLength; i++ {
		v := 0
		if len(nums1) == 0 {
			v = nums2[0]
			nums2 = nums2[1:]
		} else if len(nums2) == 0 {
			v = nums1[0]
			nums1 = nums1[1:]
		} else {
			if nums1[0] < nums2[0] {
				v = nums1[0]
				nums1 = nums1[1:]
			} else {
				v = nums2[0]
				nums2 = nums2[1:]
			}
		}

		switch i {
		case p2:
			sum += v
			break Loop
		case p1:
			sum += v
		}
	}

	if p1 == p2 {
		return float64(sum)
	}
	return float64(sum) / float64(2)
}
