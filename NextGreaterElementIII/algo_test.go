package algo

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_int2Digits(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{123}, []int{1, 2, 3}},
		{"t2", args{6792}, []int{6, 7, 9, 2}},
		{"t3", args{1043}, []int{1, 0, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := int2Digits(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("int2Digits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_int2Digits(b *testing.B) {
	n := rand.Int()
	for i := 0; i < b.N; i++ {
		digits := int2Digits(n)
		if len(digits) == 0 {
			b.Errorf("length of digits is 0")
		}
	}
}

func Test_isDecreased(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{[]int{1, 2, 3}}, false},
		{"t2", args{nil}, false},
		{"t3", args{[]int{}}, false},
		{"t4", args{[]int{5, 4, 3}}, true},
		{"t5", args{[]int{5, 4, 4, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDecreased(tt.args.nums); got != tt.want {
				t.Errorf("isDecreased() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digits2Int(t *testing.T) {
	type args struct {
		d []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 2, 3}}, 123},
		{"t2", args{[]int{0, 2, 3}}, 23},
		{"t3", args{[]int{3, 5, 0, 4}}, 3504},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digits2Int(tt.args.d); got != tt.want {
				t.Errorf("digits2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{12}, 21},
		{"t2", args{21}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.n); got != tt.want {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
