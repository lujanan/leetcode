package algorithm

// 对角线遍历
// https://leetcode-cn.com/problems/diagonal-traverse/
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) <= 0 {
		return nil
	}
	var (
		m, n            = len(matrix), len(matrix[0])
		r, c, tmpLenght int
		direction       = false
		result, tmp     []int
	)

	for i := 0; i < n+m-1; i++ {
		if i < n { // 遍历上半区对角线
			r, c = 0, i
		} else { // 遍历下半区对角线
			r, c = i-n+1, n-1
		}

		for {
			tmp = append(tmp, matrix[r][c])
			if c <= 0 || r >= m-1 {
				if !direction {
					tmpLenght = len(tmp)
					for idx := 0; idx < len(tmp)/2; idx++ {
						tmp[idx], tmp[tmpLenght-idx-1] = tmp[tmpLenght-idx-1], tmp[idx]
					}
				}
				result = append(result, tmp...)
				tmp = tmp[:0]
				direction = !direction
				break
			}
			r++
			c--
		}
	}
	return result
}
