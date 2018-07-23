package algo

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{[]int{1, 2, 3, 1}, 3, 0}, true},
		{"t2", args{[]int{1, 0, 1, 1}, 1, 2}, true},
		{"t3", args{[]int{1, 5, 9, 1, 5, 9}, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
