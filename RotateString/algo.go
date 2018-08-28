// https://leetcode.com/problems/rotate-string/description/
package algo

func rotateString(A string, B string) bool {
	if A == B {
		return true
	}

	if len(A) != len(B) {
		return false
	}

	a := A
	for i := 0; i < len(A); i++ {
		if a == B {
			return true
		}

		a = leftRotate(a)
	}

	return false
}

func leftRotate(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)

	return string(append(r[1:], r[0]))
}
