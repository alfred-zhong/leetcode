package algo

import (
	"testing"
)

func Test_findInsert(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{61, 121}, 180}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInsert(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("findInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeToInt(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"00:00"}, 0},
		{"t2", args{"01:01"}, 61},
		{"t3", args{"13:23"}, 803},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeToInt(tt.args.time); got != tt.want {
				t.Errorf("timeToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinDifference(t *testing.T) {
	type args struct {
		timePoints []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]string{"23:59", "00:00"}}, 1},
		{"t2", args{[]string{"00:00", "01:02", "01:01"}}, 1},
		{"t3", args{[]string{"05:31", "22:08", "00:35"}}, 147},
		{"t4", args{[]string{"01:01", "02:01", "03:00"}}, 59},
		{"t4", args{[]string{"01:39", "01:39", "07:20", "11:54", "00:46"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinDifference(tt.args.timePoints); got != tt.want {
				t.Errorf("findMinDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
