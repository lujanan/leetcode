// 3363.最多可收集的水果数目

package algorithm_3300

func maxCollectedFruits(fruits [][]int) int {
	var n = len(fruits)
	// (0,0)坐标的只能走固定对角线
	var dp1 int
	for i, j := 0, 0; i < n && j < n; i, j = i+1, j+1 {
		dp1 += fruits[i][j]
	}

	// (0, n-1)坐标走纵向线，最后1步到达(i,j)收益为0
	var dp2 = []int{fruits[0][n-1], 0}
	for i := 1; i < n-1; i++ {
		if i != n-2 {
			dp2[0], dp2[1] = max(dp2[0], dp2[1])+fruits[i][n-1], max(dp2[0], dp2[1])+fruits[i][n-2]
		} else {
			dp2[0], dp2[1] = max(dp2[0], dp2[1])+fruits[i][n-1], max(dp2[0], dp2[1])
		}
	}

	// (n-1, 0)坐标走水平线，最后1步到达(i,j)收益为0
	var dp3 = []int{fruits[n-1][0], 0}
	for j := 1; j < n-1; j++ {
		if j != n-2 {
			dp3[0], dp3[1] = max(dp3[0], dp3[1])+fruits[n-1][j], max(dp3[0], dp3[1])+fruits[n-2][j]
		} else {
			dp3[0], dp3[1] = max(dp3[0], dp3[1])+fruits[n-1][j], max(dp3[0], dp3[1])
		}
	}

	return dp1 + max(dp2[0], dp2[1]) + max(dp3[0], dp3[1])
}
