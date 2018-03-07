package algo

import (
	"testing"
)

func Test_pow10(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{0}, 1},
		{"t1", args{1}, 10},
		{"t1", args{2}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow10(tt.args.n); got != tt.want {
				t.Errorf("pow10() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lenInt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{3}, 1},
		{"t2", args{123}, 3},
		{"t3", args{1000}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenInt(tt.args.x); got != tt.want {
				t.Errorf("lenInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{123}, 321},
		{"t2", args{120}, 21},
		{"t3", args{-123}, -321},
		{"t4", args{1534236469}, 0},
		{"t5", args{1563847412}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
