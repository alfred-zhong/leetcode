// https://blog.csdn.net/cszhouwei/article/details/33741731
package algo

import (
	"unicode/utf8"
)

func reverseString(s string) string {
	bb := make([]byte, len(s))
	i := 0
	for s != "" {
		r, n := utf8.DecodeLastRuneInString(s)
		utf8.EncodeRune(bb[i:], r)
		i += n
		s = s[:len(s)-n]
	}
	return string(bb)
}
