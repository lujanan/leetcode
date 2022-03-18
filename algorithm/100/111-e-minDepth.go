package algorithm_100

// 111. 二叉树的最小深度
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	var num int
	var fn func(root *TreeNode, n int)
	fn = func(root *TreeNode, n int) {
		if root == nil {
			return
		}
		n++
		if root.Left == nil && root.Right == nil {
			if num <= 0 || num > n {
				num = n
			}
			return
		}
		if root.Left != nil {
			fn(root.Left, n)
		}
		if root.Right != nil {
			fn(root.Right, n)
		}
	}
	fn(root, 0)
	return num
}
