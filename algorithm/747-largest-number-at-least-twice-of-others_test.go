package algorithm

import "testing"

func Test_dominantIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{3, 6, 1, 0}}, 1},
		{"t2", args{[]int{1, 2, 3, 4}}, -1},
		{"t3", args{[]int{1, 2, 3, 4}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dominantIndex(tt.args.nums); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
