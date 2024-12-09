//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
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
// Related Topics 数组 回溯 👍 2210 👎 0

package algorithm_0

import (
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var (
		pie  = make(map[int]int)
		na   = make(map[int]int)
		col  = make(map[int]int)
		nums [][]int
	)

	var fn func(row int, num []int)
	fn = func(row int, num []int) {
		if row >= n {
			nums = append(nums, num)
			return
		}

		for i := 0; i < n; i++ {
			if _, ok := col[i]; ok {
				continue
			}
			if _, ok := pie[i-row]; ok {
				continue
			}
			if _, ok := na[i+row]; ok {
				continue
			}

			col[i] = 0
			pie[i-row] = 0
			na[i+row] = 0

			fn(row+1, append(append([]int{}, num...), i))

			delete(col, i)
			delete(pie, i-row)
			delete(na, i+row)
		}
	}

	fn(0, nil)
	return format(nums, n)
}

// leetcode submit region end(Prohibit modification and deletion)
func solveNQueensV2(n int) [][]string {
	var res [][]int
	var col, pie, na = make(map[int]int), make(map[int]int), make(map[int]int)
	var fn func(row int, sub []int)
	fn = func(row int, sub []int) {
		if row >= n {
			res = append(res, sub)
			return
		}

		for i := 0; i < n; i++ {
			_, okc := col[i]
			_, okx := pie[i+row]
			_, oky := na[row-i]
			if okc || okx || oky {
				continue
			}

			col[i] = 0
			pie[i+row] = 0
			na[row-i] = 0
			fn(row+1, append(append([]int{}, sub...), i))

			delete(col, i)
			delete(pie, i+row)
			delete(na, row-i)
		}
	}
	fn(0, nil)
	return format(res, n)
}

func format(nums [][]int, n int) [][]string {
	var res [][]string
	for _, num := range nums {
		var tmp []string
		for _, v := range num {
			tmp = append(tmp, strings.Repeat(".", v)+"Q"+strings.Repeat(".", n-v-1))
		}
		res = append(res, tmp)
	}
	return res
}
