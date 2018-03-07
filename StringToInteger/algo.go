package algo

const intSize = 32 << (^uint(0) >> 63)
const max = 1<<31 - 1
const min = -(1 << 31)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}

	for str[0] == ' ' {
		str = str[1:]
	}

	sLen := len(str)
	if intSize == 32 && (0 < sLen && sLen < 10) ||
		intSize == 64 && (0 < sLen && sLen < 19) {
		// Fast path for small integers that fit int type.
		s0 := str
		if str[0] == '-' || str[0] == '+' {
			str = str[1:]
			if len(str) < 1 {
				return 0
			}
		}

		n := 0
		for _, ch := range []byte(str) {
			ch -= '0'
			if ch > 9 {
				break
			}
			n = n*10 + int(ch)
		}
		if s0[0] == '-' {
			n = -n
		}

		if n > max {
			return max
		} else if n < min {
			return min
		}

		return n
	}

	if str[0] == '-' {
		return min
	}
	return max
}
