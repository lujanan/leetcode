// 2787.将一个数字表示成幂的和的方案数

package algorithm_2700

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		n int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{10, 2}, 1},
		{"t2", args{4, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
