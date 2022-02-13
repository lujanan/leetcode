package algorithm_1000

// https://leetcode-cn.com/problems/number-of-enclaves/

func numEnclaves(grid [][]int) (res int) {
	var dx = []int{-1, 1, 0, 0}
	var dy = []int{0, 0, -1, 1}

	var fn func(i, j int)
	fn = func(i, j int) {
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0

		for k := 0; k < len(dx); k++ {
			var x, y = j + dx[k], i + dy[k]
			if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || grid[y][x] == 0 {
				continue
			}
			fn(y, x)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if (i != 0 && i != len(grid)-1 && j != 0 && j != len(grid[0])-1) || grid[i][j] == 0 {
				continue
			}
			fn(i, j)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res += grid[i][j]
		}
	}
	return
}
