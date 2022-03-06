package algorithm_400

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{2, 3}},
		{"t2", args{[]int{1, 1, 2}}, []int{1}},
		{"t3", args{[]int{1}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
