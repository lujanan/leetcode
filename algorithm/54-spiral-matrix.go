package algorithm

import "fmt"

// 螺旋矩阵
// https://leetcode-cn.com/problems/spiral-matrix/
// tips:
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
func spiralOrder(matrix [][]int) []int {
	var m = len(matrix)
	if m < 1 || m > 10 {
		return nil
	}
	var n = len(matrix[0])
	if n < 1 || n > 10 {
		return nil
	}

	var checkMap = make(map[string]bool)
	var res []int
	// 方向调整，右下左上，i、j 增加的值
	var hi, hj = []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	var i, j, h int

	for a := 0; a < m*n; a++ {
		res = append(res, matrix[i][j])
		checkMap[fmt.Sprintf("%d_%d", i, j)] = false

		if i+hi[h] >= m || i+hi[h] < 0 || j+hj[h] >= n || j+hj[h] < 0 {
			h = (h + 1) % 4
		} else if _, ok := checkMap[fmt.Sprintf("%d_%d", i+hi[h], j)]; ok && hi[h] != 0 {
			h = (h + 1) % 4
		} else if _, ok := checkMap[fmt.Sprintf("%d_%d", i, j+hj[h])]; ok && hj[h] != 0 {
			h = (h + 1) % 4
		}

		i = i + hi[h]
		j = j + hj[h]
	}
	return res
}
