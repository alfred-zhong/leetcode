package algo

func reverse(x int) int {
	r := 0
	negative := false

	if x < 0 {
		negative = true
		x = -x
	}

	l := lenInt(x)

	for x > 0 {
		p := x % 10
		r += p * pow10(l-1)

		x = x / 10
		l--
	}

	if r > (2<<30 - 1) {
		return 0
	}

	if negative {
		r = -r
	}

	return r
}

func lenInt(x int) int {
	l := 0
	r := 1
	for x >= r {
		r *= 10
		l++
	}

	return l
}

var pm = make(map[int]int)

func pow10(n int) int {
	if v, ok := pm[n]; ok {
		return v
	}

	r := 1
	for i := 0; i < n; i++ {
		pm[i] = r
		r *= 10
	}
	return r
}
