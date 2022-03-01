//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充
//。
//
//
//
//
// 示例 1：
//
//
//输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O",
//"X","X"]]
//输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都
//会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
//
//
// 示例 2：
//
//
//输入：board = [["X"]]
//输出：[["X"]]
//
//
//
//
// 提示：
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 200
// board[i][j] 为 'X' 或 'O'
//
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 733 👎 0

package algorithm_100

func solve(board [][]byte) [][]byte {
	var d = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var h, w = len(board), len(board[0])

	var fn func(y, x int) int
	fn = func(y, x int) int {
		if board[y][x] == 'X' {
			return 0
		}
		if (x == 0 || x+1 == w || y == 0 || y+1 == h) && board[y][x] == 'O' {
			return 1
		}
		board[y][x] = 'X'
		var num int
		for _, v := range d {
			var ny, nx = y + v[0], x + v[1]
			if nx >= 0 && nx < w && y >= 0 && y < h {
				num += fn(ny, nx)
			}
		}
		if num > 0 {
			board[y][x] = 'O'
		}
		return num
	}
	for i := 1; i < h-1; i++ {
		for j := 1; j < w-1; j++ {
			fn(i, j)
		}
	}
	return board
}
