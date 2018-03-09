package algo

import "testing"

func Test_reachNumber(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{3}, 2},
		{"t2", args{4}, 3},
		{"t3", args{7}, 5},
		{"t4", args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachNumber(tt.args.target); got != tt.want {
				t.Errorf("reachNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_reachNumber(b *testing.B) {
	count := 100000
	for i := 0; i < b.N; i++ {
		for i := 0; i < count; i++ {
			reachNumber(i)
		}
	}
}
