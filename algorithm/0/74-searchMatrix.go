//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -10⁴ <= matrix[i][j], target <= 10⁴
//
// Related Topics 数组 二分查找 矩阵 👍 576 👎 0

package algorithm_0

func searchMatrix(matrix [][]int, target int) bool {
	var w = len(matrix[0])
	var l, r = 0, len(matrix)*w - 1

	for l <= r {
		var m = l + (r-l)>>1
		if target == matrix[m/w][m%w] {
			return true
		}

		if target < matrix[m/w][m%w] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
