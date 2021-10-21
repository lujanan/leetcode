package algorithm_0

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	fmt.Println(float64(float64(1)/float64(3)))
	return

	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"t2", args{[]int{4, 3, 2, 1}}, []int{4, 3, 2, 2}},
		{"t3", args{[]int{9, 9, 9, 9}}, []int{1, 0, 0, 0, 0}},
		{"t4", args{[]int{9}}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
