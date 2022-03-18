package algorithm_900

// 980. 不同路径 III
// https://leetcode-cn.com/problems/unique-paths-iii/

// 回溯
func uniquePathsIII(grid [][]int) int {
	var dxy = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var h, w = len(grid), len(grid[0])
	var n = h * w
	var sy, sx, ey, ex = -1, -1, -1, -1
	var visited = make(map[[2]int]int)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 1 {
				sy, sx = i, j
			} else if grid[i][j] == 2 {
				ey, ex = i, j
			} else if grid[i][j] == -1 {
				n--
			}
		}
	}
	// 起点或终点不存在
	if sy < 0 || sy >= h || sx < 0 || sx >= w ||
		ey < 0 || ey >= h || ex < 0 || ex >= w {
		return 0
	}

	var res int
	var dfs func(y, x, step int)
	dfs = func(y, x, step int) {
		if step+1 >= n {
			if y == ey && x == ex {
				res++
			}
			return
		}

		visited[[2]int{y, x}] = 1
		step++
		for i := range dxy {
			var ny, nx = y + dxy[i][0], x + dxy[i][1]
			if ny < 0 || ny >= h || nx < 0 || nx >= w || grid[ny][nx] == -1 {
				continue
			}
			if _, ok := visited[[2]int{ny, nx}]; ok && visited[[2]int{ny, nx}] != 0 {
				continue
			}
			dfs(ny, nx, step)
		}
		visited[[2]int{y, x}] = 0
		step--
	}
	dfs(sy, sx, 0)
	return res
}
