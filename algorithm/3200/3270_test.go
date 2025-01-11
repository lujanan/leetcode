// 3270. 求出数字答案

package algorithm_3200

import "testing"

func Test_generateKey(t *testing.T) {
	type args struct {
		num1 int
		num2 int
		num3 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{1, 10, 100}, 0},
		{"t2", args{987, 879, 789}, 777},
		{"t3", args{1, 2, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateKey(tt.args.num1, tt.args.num2, tt.args.num3); got != tt.want {
				t.Errorf("generateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
