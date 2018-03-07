package lswrc

import (
	"math/rand"
	"strings"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"abcabcbb"}, 3},
		{"t2", args{"bbbbb"}, 1},
		{"t3", args{"pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	len := 10000

	var sb strings.Builder
	for i := 0; i < len; i++ {
		sb.WriteByte(byte('a') + byte(rand.Intn('z'-'a')))
	}
	s := sb.String()

	for i := 0; i < b.N; i++ {
		n := lengthOfLongestSubstring(s)
		if n < 1 {
			b.Errorf("length of longest substring %d not illegal", n)
		}
	}
}
