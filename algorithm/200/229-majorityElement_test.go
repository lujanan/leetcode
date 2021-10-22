//给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
//
//
//
//
// 示例 1：
//
//
//输入：[3,2,3]
//输出：[3]
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：[1]
//
//
// 示例 3：
//
//
//输入：[1,1,1,3,3,2,2,2]
//输出：[1,2]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
// Related Topics 数组 哈希表 计数 排序 👍 465 👎 0

package algorithm_200

import (
	"reflect"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{[]int{3, 2, 3}}, []int{3}},
		{"t2", args{[]int{1}}, []int{1}},
		{"t3", args{[]int{1, 1, 1, 3, 3, 2, 2, 2}}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := majorityElement(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("majorityElement() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
