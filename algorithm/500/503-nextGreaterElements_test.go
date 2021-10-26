//给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第
//一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
//
// 示例 1:
//
//
//输入: [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数；
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//
//
// 注意: 输入数组的长度不会超过 10000。
// Related Topics 栈 数组 单调栈 👍 503 👎 0

package algorithm_500

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{[]int{1, 2, 1}}, []int{2, -1, 2}},
		{"t2", args{[]int{1, 2, 1, 4, 1, 3, 8, 5}}, []int{2, 4, 4, 8, 3, 8, -1, 8}},
		{"t3", args{[]int{1, 2, 1, 23, 2354, 7, 9, 12, 67, 12, 5, 1, 4, 8, 12, 34, 78, 4, 90}}, []int{2, 23, 23, 2354, -1, 9, 12, 67, 78, 34, 8, 4, 8, 12, 34, 78, 90, 90, 2354}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("nextGreaterElements() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
