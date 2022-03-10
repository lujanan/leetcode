package algorithm_0

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{[]int{1, 2, 3}}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"t2", args{[]int{1, 1, 2}}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		// {"t3", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
