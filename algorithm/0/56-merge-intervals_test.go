package algorithm_0

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{
			intervals: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			}},
			[][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			}},
		{"t2", args{
			intervals: [][]int{
				{1, 4},
				{4, 5},
			}},
			[][]int{
				{1, 5},
			}},
		{"t3", args{
			intervals: [][]int{
				{2, 3},
				{5, 5},
				{2, 2},
				{3, 4},
				{3, 4},
				{2, 2},
			}},
			[][]int{
				{2, 4},
				{5, 5},
			}},
		{"t4", args{
			intervals: [][]int{
				{1, 4},
				{0, 0},
			}},
			[][]int{
				{0, 0},
				{1, 4},
			}},
		{"t5", args{
			intervals: [][]int{
				{2, 3},
				{2, 2},
				{3, 3},
				{1, 3},
				{5, 7},
				{2, 2},
				{4, 6},
			}},
			[][]int{
				{1, 3},
				{4, 7},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
