package algorith_3400

import "testing"

func Test_maxSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 2, 3, 4, 5}}, 15},
		{"t2", args{[]int{1, 1, 0, 1, 1}}, 1},
		{"t3", args{[]int{1, 2, -1, -2, 1, 0, -1}}, 3},
		{"t4", args{[]int{-1, -2, -3, -4, -5}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.nums); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
