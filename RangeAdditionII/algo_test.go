package algo

import (
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		m int
		n int
		x int
		y int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"t1", args{3, 3, 2, 2}, 2, 2},
		{"t2", args{2, 2, 3, 3}, 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := intersect(tt.args.m, tt.args.n, tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("intersect() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("intersect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_maxCount(t *testing.T) {
	type args struct {
		m   int
		n   int
		ops [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{3, 3, [][]int{[]int{2, 2}, []int{3, 3}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.args.m, tt.args.n, tt.args.ops); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
