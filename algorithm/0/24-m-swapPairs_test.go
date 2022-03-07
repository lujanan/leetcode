//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
// Related Topics 递归 链表 👍 1264 👎 0

package algorithm_0

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{1, 2, 3, 4}}, []int{2, 1, 4, 3}},
		{"t2", args{[]int{1}}, []int{1}},
		{"t3", args{nil}, nil},
		{"t4", args{[]int{1, 3, 4, 6, 7, 2, 3, 4, 5, 6}}, []int{3, 1, 6, 4, 2, 7, 4, 3, 6, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l1 = new(ListNode)
			var tl1 = l1
			for _, v := range tt.args.head {
				tl1.Next = &ListNode{Val: v}
				tl1 = tl1.Next
			}

			list := swapPairs(l1.Next)
			var res []int
			for list != nil {
				res = append(res, list.Val)
				list = list.Next
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", res, tt.want)
			}
			return

			if got := swapPairs(l1.Next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
