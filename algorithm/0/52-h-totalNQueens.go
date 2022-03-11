//n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
//
// Related Topics 回溯 👍 338 👎 0

package algorithm_0

func totalNQueens(n int) int {
	var res int
	var col, pi, na = make([]int, n), make(map[int]int), make(map[int]int) // pi = x+y, na = x-y

	var dfs func(row int)
	dfs = func(row int) {
		if row >= n {
			res++
			return
		}

		for x := 0; x < n; x++ {
			if col[x] == 1 {
				continue
			}
			if _, ok := pi[row+x]; ok && pi[row+x] == 1 {
				continue
			}
			if _, ok := na[x-row]; ok && na[x-row] == 1 {
				continue
			}

			col[x], pi[row+x], na[x-row] = 1, 1, 1
			dfs(row + 1)
			col[x], pi[row+x], na[x-row] = 0, 0, 0
		}
	}

	dfs(0)
	return res
}
