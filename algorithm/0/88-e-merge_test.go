package algorithm_0

import (
	"reflect"
	"testing"
)

func Test_merge88(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, []int{1, 2, 2, 3, 5, 6}},
		{"t2", args{[]int{1}, 1, []int{}, 0}, []int{1}},
		{"t3", args{[]int{0}, 0, []int{1}, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge88(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge88() = %v, want %v", got, tt.want)
			}
		})
	}
}
