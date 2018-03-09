package algo

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}}, 2},
		{"t2", args{[]int{3, 3, 7, 7, 10, 11, 11}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_singleNonDuplicate1(b *testing.B) {
	single := 1
	ns := 10000
	nums := make([]int, 0, 2*ns-1)
	for i := 1; i <= ns; i++ {
		nums = append(nums, i)

		if i != single {
			nums = append(nums, i)
		}
	}
	for i := 0; i < b.N; i++ {
		r := singleNonDuplicate(nums)
		if r != single {
			b.Errorf("got %d, want %d", r, single)
		}
	}
}

func Benchmark_singleNonDuplicate2(b *testing.B) {
	single := 10000
	ns := 10000
	nums := make([]int, 0, 2*ns-1)
	for i := 1; i <= ns; i++ {
		nums = append(nums, i)

		if i != single {
			nums = append(nums, i)
		}
	}
	for i := 0; i < b.N; i++ {
		r := singleNonDuplicate(nums)
		if r != single {
			b.Errorf("got %d, want %d", r, single)
		}
	}
}
