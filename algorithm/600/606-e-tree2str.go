package algorithm_600

import "strconv"

// 606. 根据二叉树创建字符串
// https://leetcode-cn.com/problems/construct-string-from-binary-tree/

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

func tree2str(root *TreeNode) string {
	var res string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		res += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			return
		}

		res += "("
		if node.Left != nil {
			dfs(node.Left)
		}
		res += ")"

		if node.Right != nil {
			res += "("
			dfs(node.Right)
			res += ")"
		}
	}
	dfs(root)
	return res
}
