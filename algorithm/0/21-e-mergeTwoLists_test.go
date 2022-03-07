//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
// Related Topics 递归 链表 👍 2239 👎 0

package algorithm_0

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{1, 2, 4}, []int{1, 3, 4}}, []int{1, 1, 2, 3, 4, 4}},
		{"t2", args{nil, nil}, nil},
		{"t3", args{[]int{1, 5, 7}, []int{2, 4, 8}}, []int{1, 2, 4, 5, 7, 8}},
		{"t4", args{nil, []int{0}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l1, l2 = new(ListNode), new(ListNode)
			var tl1, tl2 = l1, l2
			for _, v := range tt.args.list1 {
				tl1.Next = &ListNode{Val: v}
				tl1 = tl1.Next
			}
			for _, v := range tt.args.list2 {
				tl2.Next = &ListNode{Val: v}
				tl2 = tl2.Next
			}

			list := mergeTwoLists(l1.Next, l2.Next)
			var res []int
			for list != nil {
				res = append(res, list.Val)
				list = list.Next
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", res, tt.want)
			}
			return

			if got := mergeTwoLists(l1.Next, l2.Next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
