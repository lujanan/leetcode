//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 5
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 20
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
// 1 <= n, m <= 300
// -10⁹ <= matrix[i][j] <= 10⁹
// 每行的所有元素从左到右升序排列
// 每列的所有元素从上到下升序排列
// -10⁹ <= target <= 10⁹
//
// Related Topics 数组 二分查找 分治 矩阵 👍 813 👎 0

package algorithm_200

// 利用原有排序，
// 从 matrix[0][n-1] , target 大，则跳到下一行最后一个数； 
// target 小于该数，则 target 在该行，再遍历；
// 遍历行的时候，如果 target < matrix[m][n] ，那 target < matrix[m+1][n] ，
// 即遍历下一行时，第 n 列及大于 n 列的数都大于 target
func searchMatrix(matrix [][]int, target int) bool {
	var x, y = len(matrix[0]) - 1, 0
	for x >= 0 && y < len(matrix) {
		if matrix[y][x] == target {
			return true
		}
		if matrix[y][x] < target {
			y++
		} else {
			x--
		}
	}
	return false
}
