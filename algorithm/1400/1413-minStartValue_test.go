package algorithm_1400

import "testing"

func Test_minStartValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{[]int{-3, 2, -3, 4, 2}}, 5},
		{"t2", args{[]int{1, 2}}, 1},
		{"t3", args{[]int{1, -2, -3}}, 5},
		{"t4", args{[]int{1, -2, -3, 1, -23, 10, 32, -56}}, 41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := minStartValue(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("minStartValue() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
