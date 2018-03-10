package algo

// https://leetcode.com/problems/nth-digit/description/
func findNthDigit(n int) int {
	for i := len(box) - 1; i >= 0; i-- {
		if n > box[i] {
			// return num % pow10(l-n+1) / pow10(l-n)
			shift := n - box[i]
			mod := shift % (i + 1)
			num := (pow10(i) - 1) + (shift / (i + 1))
			if mod != 0 {
				num++
			}

			if mod == 0 {
				return num % 10
			} else {
				return num / pow10(lengthOfNum(num)-mod) % 10
			}
		}
	}
	return 0
}

func lengthOfNum(num int) int {
	for i := 9; i >= 0; i-- {
		if num >= pow10(i) {
			return i + 1
		}
	}
	return 0
}

func pow10(n int) int {
	return pm[n]
}

var pm = make([]int, 10)
var box = make([]int, 9)

func init() {
	pp := 1
	for i := 0; i < 10; i++ {
		pm[i] = pp
		pp *= 10
	}

	r := 0
	for i := 1; i <= 9; i++ {
		box[i-1] = r
		r += (pow10(i) - pow10(i-1)) * i
	}
}
