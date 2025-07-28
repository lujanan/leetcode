// 92.反转链表II

package algorithm_0

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var ans, cut, cur, start, end *ListNode
	for i := 1; i <= right; i++ {
		cur, head = head, head.Next
		if i < left {
			if i == 1 {
				ans = cur
			}
			cut = cur

		} else if i == left {
			start, end = cur, cur
			
		} else if i == right {
			end.Next = head
			cur.Next = start
			start = cur
			if cut != nil {
				cut.Next = start
			} else {
				ans = start
			}

		} else {
			cur.Next = start
			start = cur
		}
	}

	return ans
}
