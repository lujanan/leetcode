package algorithm

// 螺旋矩阵
// https://leetcode-cn.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 0 {
		return nil
	}

	var (
		m, n     = len(matrix), len(matrix[0])
		dr, dc   = []int{1, 0, -1, 0}, []int{0, 1, 0, -1}
		c, r, di int
		result   []int
		check    [][]bool
	)
	check = make([][]bool, m)
	for i := 0; i < m; i++ {
		check[i] = make([]bool, n)
	}

	for i := 0; i < m*n; i++ {
		result = append(result, matrix[c][r])
		check[c][r] = true

		if c+dc[di] < 0 || c+dc[di] >= m || r+dr[di] < 0 || r+dr[di] >= n || check[c+dc[di]][r+dr[di]] {
			di = (di + 1) % 4
		}
		c = c + dc[di]
		r = r + dr[di]
	}
	return result
}
