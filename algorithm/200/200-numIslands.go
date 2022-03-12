package algorithm_200

// 并查集
func numIslands(grid [][]byte) int {
	var dxy = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var h, w = len(grid), len(grid[0])
	var dp = make([]int, h*w)
	var cnt int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '1' {
				dp[i*w+j] = i*w + j
				cnt++
			}
		}
	}

	var find = func(i, j int) int {
		var root = i*w + j
		i = i*w + j
		for root != dp[root] {
			root = dp[root]
		}
		for i != dp[i] {
			i, dp[i] = dp[i], root
		}
		return root
	}

	var union func(i, j, parent int)
	union = func(i, j, parent int) {
		p1 := find(i, j)
		if p1 != parent {
			grid[i][j] = '0'
			dp[i*w+j] = parent
			cnt--
		}
		for k := 0; k < 4; k++ {
			ny, nx := i+dxy[k][0], j+dxy[k][1]
			if ny >= 0 && ny < h && nx >= 0 && nx < w && grid[ny][nx] == '1' {
				union(ny, nx, parent)
			}
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] != '1' {
				continue
			}
			union(i, j, dp[i*w+j])
		}
	}

	return cnt
}

// 深度优先搜索
func numIslands1(grid [][]byte) (res int) {
	var dx = []int{-1, 1, 0, 0}
	var dy = []int{0, 0, -1, 1}

	var fn func(i, j int) int
	fn = func(i, j int) int {
		if grid[i][j] == '0' {
			return 0
		}
		grid[i][j] = '0'
		for k := 0; k < len(dx); k++ {
			var x, y = j + dx[k], i + dy[k]
			if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) || grid[y][x] == '0' {
				continue
			}
			fn(y, x)
		}
		return 1
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res += fn(i, j)
			}
		}
	}
	return
}
