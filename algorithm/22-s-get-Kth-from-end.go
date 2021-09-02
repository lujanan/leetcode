//输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
// 例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//
//
//
// 示例：
//
//
//给定一个链表: 1->2->3->4->5, 和 k = 2.
//
//返回链表 4->5.
// Related Topics 链表 双指针
// 👍 252 👎 0

package algorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
快慢指针的思想。我们将第一个指针 textitfast  指向链表的第 k + 1 个节点，第二个指针 textitslow  指向链表的第一个节点，
此时指针 textitfast  与 textitslow  二者之间刚好间隔 k 个节点。
此时两个指针同步向后走，当第一个指针 textitfast  走到链表的尾部空节点时，
则此时 textitslow  指针刚好指向链表的倒数第k个节点。
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var fast, slow = head, head
	for k > 1 {
		k--
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func getKthFromEnd1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var length = 1
	var tmp = head
	for tmp.Next != nil {
		length++
		tmp = tmp.Next
	}
	if k <= length {
		tmp = head
		for i := 0; i < length; i++ {
			if i+k == length {
				return tmp
			} else {
				tmp = tmp.Next
			}
		}
	}
	return nil
}
