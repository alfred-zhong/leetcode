package algo

import (
	"testing"
)

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{1}, 1},
		{"t2", args{10}, 2},
		{"t3", args{9}, 1},
		{"t4", args{999}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfNum(tt.args.num); got != tt.want {
				t.Errorf("lengthOfNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
