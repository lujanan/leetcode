package algorithm_500

// 547. 省份数量
// https://leetcode-cn.com/problems/number-of-provinces/

func findCircleNum(isConnected [][]int) int {
	var n = len(isConnected)
	var cnt = n

	var parent = make([]int, n+1)
	var initSet = func() {
		for i := 1; i <= n; i++ {
			parent[i] = i
		}
	}

	var findParent = func(i int) int {
		var root = i
		for root != parent[root] {
			root = parent[root]
		}
		for i != parent[i] {
			i, parent[i] = parent[i], root
		}
		return root
	}

	var unionSet = func(i, j int) {
		pi := findParent(i)
		pj := findParent(j)

		parent[pj] = pi
		cnt--
	}

	initSet()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 && findParent(i+1) != findParent(j+1) {
				unionSet(i+1, j+1)
			}
		}
	}
	return cnt
}
