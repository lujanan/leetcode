package algorithm_0

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{[]int{1, 2, 3}}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {2, 3}, {1, 3}, {1, 2, 3}}},
		{"t2", args{[]int{0}}, [][]int{{}, {0}}},
		{"t3", args{[]int{9, 0, 3, 5, 7}}, [][]int{{}, {0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
