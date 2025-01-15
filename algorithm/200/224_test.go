// 224. 基本计算器

package algorithm_200

import "testing"

func Test_calculateV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"1 + 1"}, 2},
		{"t2", args{" 2-1 + 2 "}, 3},
		{"t3", args{"(1+(4+5+2)-3)+(6+8)"}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateV2(tt.args.s); got != tt.want {
				t.Errorf("calculateV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
