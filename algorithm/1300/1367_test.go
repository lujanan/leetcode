// 1367. 二叉树中的链表
//给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。
//
// 如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False
//。
//
// 一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。
//
//
//
// 示例 1：
//
//
//
// 输入：head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,
//null,1,3]
//输出：true
//解释：树中蓝色的节点构成了与链表对应的子路径。
//
//
// 示例 2：
//
//
//
// 输入：head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,
//null,1,3]
//输出：true
//
//
// 示例 3：
//
// 输入：head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,
//null,1,3]
//输出：false
//解释：二叉树中不存在一一对应链表的路径。
//
//
//
//
// 提示：
//
//
// 二叉树和链表中的每个节点的值都满足 1 <= node.val <= 100 。
// 链表包含的节点数目在 1 到 100 之间。
// 二叉树包含的节点数目在 1 到 2500 之间。
//
//
// Related Topics 树 深度优先搜索 链表 二叉树 👍 204 👎 0

package algorithm_1300

import "testing"

func Test_isSubPath(t *testing.T) {
	type args struct {
		head []int
		root []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{[]int{4, 2, 8}, []int{1, 4, 4, 0, 2, 2, 0, 1, 0, 6, 8, 0, 0, 0, 0, 1, 3}}, true},
		{"t2", args{[]int{1, 4, 2, 6}, []int{1, 4, 4, 0, 2, 2, 0, 1, 0, 6, 8, 0, 0, 0, 0, 1, 3}}, true},
		{"t3", args{[]int{1, 4, 2, 6, 8}, []int{1, 4, 4, 0, 2, 2, 0, 1, 0, 6, 8, 0, 0, 0, 0, 1, 3}}, false},
		{"t4", args{[]int{4, 2}, []int{4, 4, 4, 1, 0, 0, 0, 2}}, false},
	}
	for _, tt := range tests {
		var head, tmpHead *ListNode
		for i := 0; i < len(tt.args.head); i++ {
			h := &ListNode{Val: tt.args.head[i]}
			if head == nil {
				head, tmpHead = h, h
			} else {
				tmpHead.Next = h
			}
			tmpHead = h
		}

		var builTreeFn func(idx int) *TreeNode
		builTreeFn = func(idx int) *TreeNode {
			if idx >= len(tt.args.root) {
				return nil
			}
			if tt.args.root[idx] <= 0 {
				return nil
			}

			var node = &TreeNode{Val: tt.args.root[idx]}
			node.Left = builTreeFn(2*idx + 1)
			node.Right = builTreeFn(2*idx + 2)
			return node
		}
		var root = builTreeFn(0)

		t.Run(tt.name, func(t *testing.T) {
			if got := isSubPath(head, root); got != tt.want {
				t.Errorf("isSubPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
