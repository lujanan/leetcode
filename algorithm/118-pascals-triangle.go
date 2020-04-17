package algorithm

//杨辉三角
//https://leetcode-cn.com/problems/pascals-triangle/
func generate(numRows int) [][]int {
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
