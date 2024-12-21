package algorithm_300

import (
	"fmt"
	"math"
)

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

func rob(root *TreeNode) int {
	var nodes1, nodes2 = []*TreeNode{root}, []*TreeNode{}
	var levelSum int
	var dp = [2]int{0, 0}
	for len(nodes1) > 0 {
		for i := 0; i < len(nodes1); i++ {
			if nodes1[i].Left != nil {
				nodes2 = append(nodes2, nodes1[i].Left)
			}
			if nodes1[i].Right != nil {
				nodes2 = append(nodes2, nodes1[i].Right)
			}
			levelSum += nodes1[i].Val
		}
		dp[1], dp[0] = int(math.Max(float64(dp[1]), float64(dp[0]+levelSum))), dp[1]
		fmt.Println(levelSum)
		fmt.Println(dp)
		nodes1 = nodes2
		nodes2 = make([]*TreeNode, 0)
		levelSum = 0
	}

	return int(math.Max(float64(dp[0]), float64(dp[1])))
}
