//用一个大小为 m x n 的二维网格 grid 表示一个箱子。你有 n 颗球。箱子的顶部和底部都是开着的。
//
// 箱子中的每个单元格都有一个对角线挡板，跨过单元格的两个角，可以将球导向左侧或者右侧。
//
//
// 将球导向右侧的挡板跨过左上角和右下角，在网格中用 1 表示。
// 将球导向左侧的挡板跨过右上角和左下角，在网格中用 -1 表示。
//
//
// 在箱子每一列的顶端各放一颗球。每颗球都可能卡在箱子里或从底部掉出来。如果球恰好卡在两块挡板之间的 "V" 形图案，或者被一块挡导向到箱子的任意一侧边上，就
//会卡住。
//
// 返回一个大小为 n 的数组 answer ，其中 answer[i] 是球放在顶部的第 i 列后从底部掉出来的那一列对应的下标，如果球卡在盒子里，则返回
//-1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-
//1,-1]]
//输出：[1,-1,-1,-1,-1]
//解释：示例如图：
//b0 球开始放在第 0 列上，最终从箱子底部第 1 列掉出。
//b1 球开始放在第 1 列上，会卡在第 2、3 列和第 1 行之间的 "V" 形里。
//b2 球开始放在第 2 列上，会卡在第 2、3 列和第 0 行之间的 "V" 形里。
//b3 球开始放在第 3 列上，会卡在第 2、3 列和第 0 行之间的 "V" 形里。
//b4 球开始放在第 4 列上，会卡在第 2、3 列和第 1 行之间的 "V" 形里。
//
//
// 示例 2：
//
//
//输入：grid = [[-1]]
//输出：[-1]
//解释：球被卡在箱子左侧边上。
//
//
// 示例 3：
//
//
//输入：grid = [[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1],[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1]
//]
//输出：[0,1,2,3,4,-1]
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 100
// grid[i][j] 为 1 或 -1
//
// Related Topics 深度优先搜索 数组 动态规划 矩阵 模拟 👍 90 👎 0

package algorithm_1700

func findBall(grid [][]int) []int {
	var n = len(grid[0])
	var res = make([]int, n)
	for col := range res {
		j := col
		for _, row := range grid {
			dir := row[col]
			col += dir
			if col < 0 || col >= n || row[col] != dir {
				col = -1
				break
			}
		}
		res[j] = col
	}
	return res
}

// dp2
func findBall3(grid [][]int) []int {
	var n = len(grid[0])
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i
	}

	for y := 0; y < len(grid); y++ {
		for i := 0; i < n; i++ {
			if res[i] < 0 {
				continue
			}
			if grid[y][res[i]] == -1 && res[i]-1 >= 0 && grid[y][res[i]-1] == -1 {
				res[i] = res[i] - 1
			} else if grid[y][res[i]] == 1 && res[i]+1 < n && grid[y][res[i]+1] == 1 {
				res[i] = res[i] + 1
			} else {
				res[i] = -1
			}
		}
	}

	return res
}

// dp1
func findBall2(grid [][]int) []int {
	var n = len(grid[0])
	var dp = make([][]int, len(grid)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	var res = make([]int, n)
	res[0] = -1
	dp[0][0] = n
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = i
		res[i] = -1
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if dp[i-1][j] <= 0 {
				continue
			}

			nx := j + grid[i-1][j]
			if nx < 0 || nx >= n || grid[i-1][j] != grid[i-1][nx] {
				continue
			}
			dp[i][nx] = dp[i-1][j]
		}
	}
	for i := 0; i < n; i++ {
		if dp[len(dp)-1][i] == n {
			res[0] = i
		} else if dp[len(dp)-1][i] > 0 {
			res[dp[len(dp)-1][i]] = i
		}
	}
	return res
}

// 深度优先搜索
func findBall1(grid [][]int) []int {
	var m, n = len(grid), len(grid[0])
	var fn func(y, x int) int
	fn = func(y, x int) int {
		if y >= m {
			return x
		}

		var nx = x + grid[y][x]
		if nx < 0 || nx >= n || grid[y][x] != grid[y][nx] {
			return -1
		}
		return fn(y+1, nx)
	}

	var res []int
	for i := 0; i < len(grid[0]); i++ {
		res = append(res, fn(0, i))
	}
	return res
}
