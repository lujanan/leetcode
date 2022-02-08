package algorithm_0

import "math"

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

func isValidBST(root *TreeNode) bool {
	return checkSub(root, math.MinInt64, math.MaxInt64)
}

func checkSub(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}
	if root.Left != nil &&
		(root.Left.Val >= root.Val || int64(root.Left.Val) <= min || int64(root.Left.Val) >= max) {
		return false
	}
	if root.Right != nil &&
		(root.Right.Val <= root.Val || int64(root.Right.Val) <= min || int64(root.Right.Val) >= max) {
		return false
	}

	if !checkSub(root.Left, min, int64(root.Val)) {
		return false
	}
	return checkSub(root.Right, int64(root.Val), max)
}
