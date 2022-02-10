package algorithm_200

func numIslands(grid [][]byte) (res int) {
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
