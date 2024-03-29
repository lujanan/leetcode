//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
// 问总共有多少条不同的路径？
//
//
//
// 示例 1：
//
//
//输入：m = 3, n = 7
//输出：28
//
// 示例 2：
//
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//
//
// 示例 3：
//
//
//输入：m = 7, n = 3
//输出：28
//
//
// 示例 4：
//
//
//输入：m = 3, n = 3
//输出：6
//
//
//
// 提示：
//
//
// 1 <= m, n <= 100
// 题目数据保证答案小于等于 2 * 10⁹
//
// Related Topics 数学 动态规划 组合数学 👍 1279 👎 0

package algorithm_0

// 动态规划
func uniquePaths(m int, n int) int {
	var dp []int
	for i := 0; i < n; i++ {
		dp = append(dp, 1)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

// bfs
func uniquePaths1(m int, n int) int {
	var dyx = [][]int{{0, 1}, {1, 0}}
	var curr = map[[2]int]int64{{0, 0}: 1}
	var next = make(map[[2]int]int64)
	var dp = map[[2]int]int64{{0, 0}: 1}

	for len(curr) > 0 {
		for point, step := range curr {
			for k := 0; k < 2; k++ {
				var ny, nx = point[0] + dyx[k][0], point[1] + dyx[k][1]
				if ny < 0 || ny >= m || nx < 0 || nx >= n {
					continue
				}
				dp[[2]int{ny, nx}] += step
				if !(ny+1 == m && nx+1 == n) {
					next[[2]int{ny, nx}] = dp[[2]int{ny, nx}]
				}
			}
		}
		curr = next
		next = make(map[[2]int]int64)
	}
	return int(dp[[2]int{m - 1, n - 1}])
}
