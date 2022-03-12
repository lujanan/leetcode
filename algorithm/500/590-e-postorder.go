package algorithm_500

// 590. N 叉树的后序遍历
// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

//  递归
func postorder(root *Node) []int {
	var res []int

	var fn func(node *Node)
	fn = func(node *Node) {
		if node == nil {
			return
		}
		if len(node.Children) < 1 {
			res = append(res, node.Val)
			return
		}

		for _, v := range node.Children {
			fn(v)
		}
		res = append(res, node.Val)
	}
	fn(root)
	return res
}
