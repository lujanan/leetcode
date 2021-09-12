package algorithm_200

// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

// 提示：

// 列表中的节点数目在范围 [0, 10^4] 内
// 1 <= Node.val <= 50
// 0 <= val <= 50

// 示例1：
// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]

// 示例2：
// 输入：head = [], val = 1
// 输出：[]

// 示例3;
// 输入：head = [7,7,7,7], val = 7
// 输出：[]

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Val == val {
			head = head.Next
			continue
		}
		break
	}
	var tmp = head
	for tmp != nil {
		if tmp.Next != nil && tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
			continue
		}
		tmp = tmp.Next
	}
	return head
}
