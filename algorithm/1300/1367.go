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

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
    if root == nil {
        return false
    }
    return dfs(root, head) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(rt *TreeNode, head *ListNode) bool {
    // 链表已经全部匹配完，匹配成功
    if head == nil {
        return true
    }
    // 二叉树访问到了空节点，匹配失败
    if rt == nil {
        return false
    }
    // 当前匹配的二叉树上节点的值与链表节点的值不相等，匹配失败
    if rt.Val != head.Val {
        return false
    }
    return dfs(rt.Left, head.Next) || dfs(rt.Right, head.Next)
}


func isSubPathV2(head *ListNode, root *TreeNode) bool {
	var headArr []int
	for head != nil {
		headArr = append(headArr, head.Val)
		head = head.Next
	}

	var rootList = make([][]int, 0)
	var dfs func(node *TreeNode, path []int)
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			rootList = append(rootList, path)
			return
		}

		path = append(path, node.Val)
		dfs(node.Left, append([]int{}, path...))
		dfs(node.Right, append([]int{}, path...))
	}
	dfs(root, []int{})

	var isSubArr = func(arr, subArr []int) bool {
		if len(arr) < len(subArr) {
			return false
		}

		var found bool
		for i := 0; i <= len(arr)-len(subArr); i++ {
			found = false
			if arr[i] == subArr[0] {
				found = true
				for j, k := 1, len(subArr)-1; k >= j; j, k = j+1, k-1 {
					if arr[i+j] != subArr[j] || arr[i+k] != subArr[k] {
						found = false
						break
					}
				}
			}
			if found {
				return true
			}
		}
		return false
	}

	for _, path := range rootList {
		if isSubArr(path, headArr) {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
