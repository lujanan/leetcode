// 221.最大正方形
//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
//
//
// 示例 1：
//
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：4
//
//
// 示例 2：
//
//
//输入：matrix = [["0","1"],["1","0"]]
//输出：1
//
//
// 示例 3：
//
//
//输入：matrix = [["0"]]
//输出：0
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] 为 '0' 或 '1'
//
//
// Related Topics 数组 动态规划 矩阵 👍 1739 👎 0

package algorithm_200

// leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	var dp = make([][2]int, len(matrix[0]))
	for j := 0; j < len(matrix[0]); j++ {
		dp[j][0], dp[j][1] = 0, int(matrix[0][j]-'0')
		if j > 0 {
			dp[j][0] = max(dp[j-1][0], dp[j-1][1])
		}
	}
	for i := 1; i < len(matrix); i++ {
		dp[0][0], dp[0][1] = max(dp[0][0], dp[0][1]), int(matrix[i][0]-'0')

		for j := 1; j < len(matrix[i]); j++ {
			dp[j][0] = max(max(dp[j-1][0], dp[j-1][1]), max(dp[j][0], dp[j][1]))
			if matrix[i][j] == '0' {
				dp[j][1] = 0
			} else {
				r := min(dp[j][1], dp[j-1][1]) // 找到包含左边或上面位置的最小正方形边长
				dp[j][1] = r + int(matrix[i-r][j-r]-'0')
				// 以当前位置为正方形右下角时，根据上面找到的边长，判断左上角(对角)位置是否为1，是则边长+1，否则边长不变
				// 如下图，matrix[3][3] 位置，左边和上面位置最小边长为2，但的左上角 matrix[1][1] 为0，该位置的正方形边长最长为2
				// {'1','0','1','0','0'},
				// {'1','0','1','1','1'},
				// {'1','1','1','1','1'},
				// {'1','1','1','1','0'}
			}
		}
	}

	l := max(dp[len(dp)-1][0], dp[len(dp)-1][1])
	return l * l
}

//leetcode submit region end(Prohibit modification and deletion)
