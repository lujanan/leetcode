// 59.螺旋矩阵II
//给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 20
//
//
// Related Topics 数组 矩阵 模拟 👍 1384 👎 0

package algorithm_0

// leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	var direction = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var d, x, y int

	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for i := 1; i <= n*n; i++ {
		res[y][x] = i
		if i >= n*n {
			break
		}

		for x+direction[d][0] < 0 || x+direction[d][0] >= n ||
			y+direction[d][1] < 0 || y+direction[d][1] >= n ||
			res[y+direction[d][1]][x+direction[d][0]] > 0 {
			d = (d + 1) % 4
		}
		x, y = x+direction[d][0], y+direction[d][1]
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
