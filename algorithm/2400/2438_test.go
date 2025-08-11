// 2438.二的幂数组中查询范围内的乘积

package algorithm_2000

import (
	"reflect"
	"testing"
)

func Test_productQueries(t *testing.T) {
	type args struct {
		n       int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{15, [][]int{{0, 1}, {2, 2}, {0, 3}}}, []int{2, 4, 64}},
		{"t2", args{2, [][]int{{0, 0}}}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productQueries(tt.args.n, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
