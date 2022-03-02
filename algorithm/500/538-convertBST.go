package algorithm_500

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	var num int
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}

		fn(node.Right)
		num += node.Val
		node.Val = num
		fn(node.Left)
	}
	fn(root)
	return root
}
