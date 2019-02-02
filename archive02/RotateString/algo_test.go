package algo

import "testing"

func Test_rotateString(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{"abcde", "cdeab"}, true},
		{"t2", args{"abcde", "abced"}, false},
		{"t3", args{"gcmbf", "fgcmb"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
