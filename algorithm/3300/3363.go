// 3363.最多可收集的水果数目

package algorithm_3300

func maxCollectedFruits(fruits [][]int) int {
	var n = len(fruits)
	// (0,0)坐标的只能走固定对角线
	var dp1 = fruits[0][0] + fruits[n-1][n-1]

	// 其余两个只能走自己所在的上下半区，不能越过对角线
	// (0, n-1)坐标走右上角上半区，最后1步到达(i,j)收益为0
	// (n-1, 0)坐标走左下角下半区，最后1步到达(i,j)收益为0
	var dp2, dp3 [2][]int
	for i := 0; i < 2; i++ {
		dp2[i] = make([]int, n-1)
		dp3[i] = make([]int, n-1)
	}
	dp2[0][0] = fruits[0][n-1]
	dp3[0][0] = fruits[n-1][0]

	for i := 1; i < n-1; i++ {
		dp1 += fruits[i][i]
		for j := 0; n-1-j > i && j <= i; j++ {
			if j > 0 {
				dp2[1][j] = max(dp2[0][j-1], dp2[0][j], dp2[0][j+1]) + fruits[i][n-1-j]
				dp3[1][j] = max(dp3[0][j-1], dp3[0][j], dp3[0][j+1]) + fruits[n-1-j][i]
			} else {
				dp2[1][j] = max(dp2[0][j], dp2[0][j+1]) + fruits[i][n-1-j]
				dp3[1][j] = max(dp3[0][j], dp3[0][j+1]) + fruits[n-1-j][i]
			}
		}
		dp2[0], dp2[1] = dp2[1], dp2[0]
		dp3[0], dp3[1] = dp3[1], dp3[0]
	}

	return dp1 + dp2[0][0] + dp3[0][0]
}
