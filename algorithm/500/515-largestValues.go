//给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
//
//
//
// 示例1：
//
//
//
//
//输入: root = [1,3,2,5,3,null,9]
//输出: [1,3,9]
//
//
// 示例2：
//
//
//输入: root = [1,2,3]
//输出: [1,3]
//
//
//
//
// 提示：
//
//
// 二叉树的节点个数的范围是 [0,10⁴]
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 168 👎 0

package algorithm_500

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}

	var cur = []*TreeNode{root}
	var next []*TreeNode

loop:
	res = append(res, math.MinInt64)
	for _, node := range cur {
		if node.Val > res[len(res)-1] {
			res[len(res)-1] = node.Val
		}
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
	}

	if len(next) > 0 {
		cur = cur[:0]
		cur = append(cur, next...)
		next = next[:0]
		goto loop
	}
	return
}
