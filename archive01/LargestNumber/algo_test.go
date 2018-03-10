package algo

import (
	"testing"
)

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{[]int{3, 30, 34, 5, 9}}, "9534330"},
		{"t2", args{[]int{128, 12}}, "12812"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimZero(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"0002"}, "2"},
		{"t2", args{"123"}, "123"},
		{"t3", args{"00"}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimZero(tt.args.s); got != tt.want {
				t.Errorf("trimZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
