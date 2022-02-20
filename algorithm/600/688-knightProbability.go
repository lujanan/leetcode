//在一个 n x n 的国际象棋棋盘上，一个骑士从单元格 (row, column) 开始，并尝试进行 k 次移动。行和列是 从 0 开始 的，所以左上单元格
//是 (0,0) ，右下单元格是 (n - 1, n - 1) 。
//
// 象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。
//
//
//
// 每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。
//
// 骑士继续移动，直到它走了 k 步或离开了棋盘。
//
// 返回 骑士在棋盘停止移动后仍留在棋盘上的概率 。
//
//
//
// 示例 1：
//
//
//输入: n = 3, k = 2, row = 0, column = 0
//输出: 0.0625
//解释: 有两步(到(1,2)，(2,1))可以让骑士留在棋盘上。
//在每一个位置上，也有两种移动可以让骑士留在棋盘上。
//骑士留在棋盘上的总概率是0.0625。
//
//
// 示例 2：
//
//
//输入: n = 1, k = 0, row = 0, column = 0
//输出: 1.00000
//
//
//
//
// 提示:
//
//
// 1 <= n <= 25
// 0 <= k <= 100
// 0 <= row, column <= n
//
// Related Topics 动态规划 👍 170 👎 0

package algorithm_600

// 题目理解：
// 1、马跳到棋盘外就不在继续
// 2、每步都要计算当前的概率
// 3、每进行一步仍留在棋盘的概率 = 当前留在棋盘的概率 x 上一步仍留在棋盘的概率
func knightProbability(n int, k int, row int, column int) float64 {
	var dy = []int{-2, -2, -1, 1, 2, 2, 1, -1}
	var dx = []int{-1, 1, 2, 2, 1, -1, -2, -2}
	var dp = make([][][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]float64, k+1)
			dp[i][j][0] = 1
		}
	}
	for i := 1; i <= k; i++ {
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				for idx := 0; idx < 8; idx++ {
					// 当前可跳的位置
					nx, ny := x+dx[idx], y+dy[idx]
					if nx < 0 || nx >= n || ny < 0 || ny >= n {
						continue
					}
					// 有8个方向可跳，每方向概率 1/8
					dp[y][x][i] += dp[ny][nx][i-1] / 8
				}
			}
		}
	}
	return dp[row][column][k]
}

//
func knightProbability1(n int, k int, row int, column int) float64 {
	var dp float64 = 1
	var dy = []int{-2, -2, -1, 1, 2, 2, 1, -1}
	var dx = []int{-1, 1, 2, 2, 1, -1, -2, -2}

	var next = [][2]int{{row, column}}
	for i := 0; i < k; i++ {
		var tmp = make([][2]int, 0)
		var in float64
		for _, node := range next {
			for k := 0; k < 8; k++ {
				nx, ny := node[1]+dx[k], node[0]+dy[k]
				if nx < 0 || nx >= n || ny < 0 || ny >= n {
					continue
				}
				in++
				tmp = append(tmp, [2]int{ny, nx})
			}
		}

		if len(next) > 0 {
			dp = dp * in / float64(len(next)) / 8
			next = tmp
		}
	}
	return dp
}
