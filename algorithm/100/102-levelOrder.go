//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 广度优先搜索 二叉树 👍 1170 👎 0

package algorithm_100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}
	var fn func(root *TreeNode, level int)
	fn = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(res) < level {
			res = append(res, []int{})
		}
		res[level-1] = append(res[level-1], root.Val)
		fn(root.Left, level+1)
		fn(root.Right, level+1)
	}
	fn(root, 1)
	return
}
