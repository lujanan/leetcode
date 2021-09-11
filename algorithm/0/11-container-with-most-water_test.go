package algorithm_0

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"t2", args{height: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 20},
		{"t3", args{height: []int{1, 3, 5, 7, 9, 11, 13, 15, 17}}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
