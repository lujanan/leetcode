package algorithm_100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue = []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		n, m := queue[0], queue[1]
		queue = queue[2:]
		if n == nil && m == nil {
			continue
		}
		if n == nil || m == nil {
			return false
		}
		if n.Val != m.Val {
			return false
		}

		queue = append(queue, n.Left, m.Right)
		queue = append(queue, n.Right, m.Left)
	}
	return true
}

// 递归
func isSymmetric0(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return check(l.Left, r.Right) && check(l.Right, r.Left)
}
