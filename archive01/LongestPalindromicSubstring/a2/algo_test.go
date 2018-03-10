package a2

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"babad"}, "aba"},
		{"t2", args{"cbbd"}, "bb"},
		{"t3", args{"qwreterbb"}, "reter"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestIndexCenter(t *testing.T) {
	type args struct {
		x int
		y int
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{2, 2, "abcba"}, "abcba"},
		{"t2", args{2, 2, "abcbd"}, "bcb"},
		{"t3", args{1, 2, "abbd"}, "bb"},
		{"t4", args{2, 2, "abcbac"}, "abcba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIndexCenter(tt.args.x, tt.args.y, tt.args.s); got != tt.want {
				t.Errorf("longestIndexCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}
