package algorithm_0

// 94. 二叉树的中序遍历
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var fn func(root *TreeNode)
	fn = func(root *TreeNode) {
		if root == nil {
			return
		}
		fn(root.Left)
		res = append(res, root.Val)
		fn(root.Right)
	}
	fn(root)

	return res
}
