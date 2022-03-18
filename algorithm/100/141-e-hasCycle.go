package algorithm_100

// 141. 环形链表
// https://leetcode-cn.com/problems/linked-list-cycle/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  解法1: hash

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法2: 快慢指针， 快指针追上慢指针，则有环存在
func hasCycle(head *ListNode) bool {
	var slow, fast = head, head
	if fast != nil {
		fast = fast.Next
	}
	var mf bool
	for slow != nil || fast != nil {
		if slow != nil && fast != nil && slow == fast {
			return true
		}
		if !mf && slow != nil {
			slow = slow.Next
		}
		mf = !mf
		if fast != nil {
			fast = fast.Next
		}
	}

	return false
}
