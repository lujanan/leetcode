//给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
//
//
//
// 示例 1：
//
//
//输入: root = [5,3,6,2,4,null,7], k = 9
//输出: true
//
//
// 示例 2：
//
//
//输入: root = [5,3,6,2,4,null,7], k = 28
//输出: false
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1, 10⁴].
// -10⁴ <= Node.val <= 10⁴
// root 为二叉搜索树
// -10⁵ <= k <= 10⁵
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 哈希表 双指针 二叉树 👍 352 👎 0

package algorithm_600

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	var nMap = make(map[int]int)
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := nMap[k-node.Val]; ok {
			return true
		}
		nMap[node.Val] = 0
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
