package algorithm_2000

import "testing"

func Test_countKDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{[]int{1, 2, 2, 1}, 1}, 4},
		{"t2", args{[]int{1, 3}, 3}, 0},
		{"t3", args{[]int{3, 2, 1, 5, 4}, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := countKDifference(tt.args.nums, tt.args.k); gotRes != tt.wantRes {
				t.Errorf("countKDifference() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
