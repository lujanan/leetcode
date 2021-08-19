//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针
// 👍 1175 👎 0

package algorithm

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
		{"t2", args{[]int{0, 1}}, []int{1, 0}},
		{"t3", args{[]int{0, 1, 1, 3, 12}}, []int{1, 1, 3, 12, 0}},
		{"t4", args{[]int{0, 0, 0, 0, 0}}, []int{0, 0, 0, 0, 0}},
		{"t5", args{[]int{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
