package twoSum

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"t2", args{[]int{3, 6, 8, 4}, 10}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_twoSum(b *testing.B) {
	size := 10000
	numMax := 100

	nums := make([]int, size)
	for i := range nums {
		nums[i] = rand.Intn(numMax)
	}
	target := rand.Intn(numMax)
	p1 := rand.Intn(numMax)
	p2 := target - p1

	i1 := rand.Intn(size)
	i2 := i1
	for i2 == i1 {
		i2 = rand.Intn(size)
	}
	nums[i1] = p1
	nums[i2] = p2

	for i := 0; i < b.N; i++ {
		r := twoSum(nums, target)
		if r == nil || len(r) != 2 {
			b.Errorf("result %v is not illegal", r)
			continue
		}
		if nums[r[0]]+nums[r[1]] != target {
			b.Errorf("%d + %d != %d", nums[r[0]], nums[r[1]], target)
		}
	}
}
