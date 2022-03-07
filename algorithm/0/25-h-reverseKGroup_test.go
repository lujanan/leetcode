//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
//
//
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
// 示例 3：
//
//
//输入：head = [1,2,3,4,5], k = 1
//输出：[1,2,3,4,5]
//
//
// 示例 4：
//
//
//输入：head = [1], k = 1
//输出：[1]
//
//
//
//
//
// 提示：
//
//
// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz
//
// Related Topics 递归 链表 👍 1504 👎 0

package algorithm_0

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{1, 2, 3, 4, 5}, 2}, []int{2, 1, 4, 3, 5}},
		{"t2", args{[]int{1, 2, 3, 4, 5}, 3}, []int{3, 2, 1, 4, 5}},
		{"t3", args{[]int{1, 2, 3, 4, 5}, 1}, []int{1, 2, 3, 4, 5}},
		{"t4", args{[]int{1, 3, 4, 6, 7, 2, 3, 4, 5, 6}, 1}, []int{1, 3, 4, 6, 7, 2, 3, 4, 5, 6}},
		{"t5", args{[]int{1}, 1}, []int{1}},
		{"t6", args{[]int{1, 2, 3, 4, 5, 34, 2, 3, 12, 3, 34, 65, 67, 23, 8, 9}, 5}, []int{5, 4, 3, 2, 1, 3, 12, 3, 2, 34, 8, 23, 67, 65, 34, 9}},
		{"t7", args{[]int{1, 2, 3, 4, 5, 34, 2, 3, 12, 3, 34, 65, 67, 23, 8, 9}, 3}, []int{3, 2, 1, 34, 5, 4, 12, 3, 2, 65, 34, 3, 8, 23, 67, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l1 = new(ListNode)
			var tl1 = l1
			for _, v := range tt.args.head {
				tl1.Next = &ListNode{Val: v}
				tl1 = tl1.Next
			}

			list := reverseKGroup(l1.Next, tt.args.k)
			var res []int
			for list != nil {
				res = append(res, list.Val)
				list = list.Next
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", res, tt.want)
			}
			return

			if got := reverseKGroup(l1.Next, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
