//给你一棵根节点为 0 的 二叉树 ，它总共有 n 个节点，节点编号为 0 到 n - 1 。同时给你一个下标从 0 开始的整数数组 parents 表示这棵
//树，其中 parents[i] 是节点 i 的父节点。由于节点 0 是根，所以 parents[0] == -1 。
//
// 一个子树的 大小 为这个子树内节点的数目。每个节点都有一个与之关联的 分数 。求出某个节点分数的方法是，将这个节点和与它相连的边全部 删除 ，剩余部分是若
//干个 非空 子树，这个节点的 分数 为所有这些子树 大小的乘积 。
//
// 请你返回有 最高得分 节点的 数目 。
//
//
//
// 示例 1:
//
//
//
// 输入：parents = [-1,2,0,2,0]
//输出：3
//解释：
//- 节点 0 的分数为：3 * 1 = 3
//- 节点 1 的分数为：4 = 4
//- 节点 2 的分数为：1 * 1 * 2 = 2
//- 节点 3 的分数为：4 = 4
//- 节点 4 的分数为：4 = 4
//最高得分为 4 ，有三个节点得分为 4 （分别是节点 1，3 和 4 ）。
//
//
// 示例 2：
//
//
//
// 输入：parents = [-1,2,0]
//输出：2
//解释：
//- 节点 0 的分数为：2 = 2
//- 节点 1 的分数为：2 = 2
//- 节点 2 的分数为：1 * 1 = 1
//最高分数为 2 ，有两个节点分数为 2 （分别为节点 0 和 1 ）。
//
//
//
//
// 提示：
//
//
// n == parents.length
// 2 <= n <= 10⁵
// parents[0] == -1
// 对于 i != 0 ，有 0 <= parents[i] <= n - 1
// parents 表示一棵二叉树。
//
// Related Topics 树 深度优先搜索 数组 二叉树 👍 39 👎 0

package algorithm_2000

func countHighestScoreNodes(parents []int) int {
	var res, n = 0, len(parents)
	var children = make([][]int, n)
	// 预处理查询每个节点的子节点
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}

	var maxScore int
	var dfs func(idx int) int
	dfs = func(idx int) int {
		var score, size = 1, n-1 // 减去当前节点剩余的节点数量(左子节点、右子节点、父节点 3部分节点之和)
		for _, v := range children[idx] {
			var cSize = dfs(v)
			score *= cSize
			// 减去所有子节点数量，则得到父节点所在部分的节点数量
			size -= cSize
		}

		// 有父节点
		if idx > 0 {
			score *= size
		}

		if maxScore == score {
			res++
		} else if maxScore < score {
			res = 1
			maxScore = score
		}
		// 返回当前节点+子节点数量
		return n - size
	}
	dfs(0)

	return res
}

func countHighestScoreNodes1(parents []int) int {
	var ll = len(parents)
	var nodes, nodeRes = make([]int, ll), make([]int, ll)
	var nodeMap = make([][]int, ll)

	var sumNode func(node int) int
	sumNode = func(idx int) int {
		if nodes[idx] <= 0 {
			if len(nodeMap[idx]) < 1 {
				nodes[idx] = 1
				return 1
			}

			var cnt = 1
			for _, v := range nodeMap[idx] {
				cnt += sumNode(v)
			}
			nodes[idx] = cnt
		}
		return nodes[idx]
	}
	for i := 1; i < ll; i++ {
		nodeMap[parents[i]] = append(nodeMap[parents[i]], i)
	}

	nodes[0], nodeRes[0] = ll, 1
	for i := 1; i < ll; i++ {
		nodeRes[i] = 1
		nodes[i] = sumNode(i)
	}

	for i := range parents {
		if i > 0 {
			nodeRes[i] *= nodes[0] - nodes[i]
		}
		if parents[i] >= 0 {
			nodeRes[parents[i]] *= nodes[i]
		}
	}

	var res, mn int
	for i := range nodeRes {
		if mn < nodeRes[i] {
			mn = nodeRes[i]
			res = 1
		} else if mn == nodeRes[i] {
			res++
		}
	}
	return res
}
