package algorithm_100

// 杨辉三角
// https://leetcode-cn.com/problems/pascals-triangle/
func generate1(numRows int) [][]int {
	var (
		result      = make([][]int, numRows)
		left, right int
	)
	for i := 0; i < numRows; i++ {
		result[i] = []int{1}
		if i > 0 {
			for n := 0; n < len(result[i-1]); n++ {
				left = result[i-1][n]
				if n+1 == len(result[i-1]) {
					right = 0
				} else {
					right = result[i-1][n+1]
				}
				result[i] = append(result[i], left+right)
			}
		}
	}
	return result
}

func generate(numRows int) [][]int {
	var ans [][]int
	for i := 0; i < numRows; i++ {
		var row = make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans = append(ans, row)
	}

	return ans
}
