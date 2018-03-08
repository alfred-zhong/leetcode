package algo

import (
	"sort"
	"reflect"
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1", args{"()())()"}, []string{"()()()", "(())()"}},
		{"t2", args{"(a)())()"}, []string{"(a)()()", "(a())()"}},
		{"t3", args{")("}, []string{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeInvalidParentheses(tt.args.s); 
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
