// 128.最长连续序列

package algorithm_100

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{100, 4, 200, 1, 3, 2}}, 4},
		{"t2", args{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}}, 9},
		{"t3", args{[]int{1, 0, 1, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
