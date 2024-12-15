package algorithm_3200

import (
	"reflect"
	"testing"
)

func Test_getFinalStateV2(t *testing.T) {
	type args struct {
		nums       []int
		k          int
		multiplier int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{2, 1, 3, 5, 6}, 5, 2}, []int{8, 4, 6, 5, 6}},
		{"t2", args{[]int{100000, 2000}, 2, 1000000}, []int{999999307, 999999993}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFinalStateV2(tt.args.nums, tt.args.k, tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFinalStateV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
