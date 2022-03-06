package algorithm_0

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t0", args{[]int{0, 1, 2}}, 3},
		{"t1", args{[]int{1, 2, 0}}, 3},
		{"t2", args{[]int{3, 4, -1, 1}}, 2},
		{"t3", args{[]int{7, 8, 9, 11, 12}}, 1},
		{"t4", args{[]int{1, 2, 10, 123, 123, 235, 346, 234, 54, 6457, 568, 345, 234, 2345, 6547, 6234, 82, 234, 657, 234, 456, 2, 4546, 423, 4436, 5567, 8234, 2323, 546, 234, 2}}, 3},
		{"t5", args{[]int{1}}, 2},
		{"t6", args{[]int{-10, -3, -100, -1000, -239, 1}}, 2},
		{"t7", args{[]int{-1}}, 1},
		{"t8", args{[]int{-5}}, 1},
		{"t9", args{[]int{0}}, 1},
		{"t10", args{[]int{1, 0, 3, 3, 0, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
