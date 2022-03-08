package algorithm_0

// 64. 最小路径和
// https://leetcode-cn.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var h, w = len(grid), len(grid[0])
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])

			} else if i-1 >= 0 && j-1 < 0 {
				grid[i][j] += grid[i-1][j]

			} else if i-1 < 0 && j-1 >= 0 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[h-1][w-1]
}

func minPathSum1(grid [][]int) int {
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var h, w = len(grid), len(grid[0])
	for i := h - 1; i >= 0; i-- {
		for j := w - 1; j >= 0; j-- {
			if i+1 < h && j+1 < w {
				grid[i][j] += min(grid[i+1][j], grid[i][j+1])
			} else if i+1 >= h && j+1 < w {
				grid[i][j] += grid[i][j+1]
			} else if i+1 < h && j+1 >= w {
				grid[i][j] += grid[i+1][j]
			}
		}
	}

	return grid[0][0]
}
