package algorithm_400

// 对角线遍历
// https://leetcode-cn.com/problems/diagonal-traverse/
func findDiagonalOrder(matrix [][]int) []int {
	var ans []int
	var direct = [][]int{{-1, 1}, {1, -1}}
	var m, n = len(matrix), len(matrix[0])
	var i, j, d int
	for k := 0; k < m*n; k++ {
		ans = append(ans, matrix[i][j])
		if 0 <= i+direct[d][0] && i+direct[d][0] < m && 0 <= j+direct[d][1] && j+direct[d][1] < n {
			i, j = i+direct[d][0], j+direct[d][1]
			continue
		}

		if d == 0 { // 斜向上
			if j+1 < n {
				j++ // 优先往右
			} else {
				i++ // 往下
			}
		} else { // 斜向下
			if i+1 < m {
				i++ // 优先往下
			} else {
				j++ // 往右
			}
		}
		d ^= 1
	}
	return ans
}

func findDiagonalOrder1(matrix [][]int) []int {
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
