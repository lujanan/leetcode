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

// 迭代
func postorder(root *Node) []int {
	var res []int
	var nodes []*Node
	if root != nil {
		nodes = []*Node{root}
	}

	for len(nodes) > 0 {
		var node = nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		if len(node.Children) > 0 {
			nodes = append(nodes, &Node{Val: node.Val})
			for i := len(node.Children) - 1; i >= 0; i-- {
				nodes = append(nodes, node.Children[i])
			}
		} else {
			res = append(res, node.Val)
		}
	}

	return res
}

//  递归
func postorder1(root *Node) []int {
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
