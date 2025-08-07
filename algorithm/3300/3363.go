// 3363.最多可收集的水果数目

package algorithm_3300

func maxCollectedFruits(fruits [][]int) int {
	var n = len(fruits)
	// (0,0)坐标的只能走固定对角线
	var dp1 int
	for i, j := 0, 0; i < n && j < n; i, j = i+1, j+1 {
		dp1 += fruits[i][j]
	}

	// 其余两个只能走自己所在的上下半区，不能越过对角线
	// (0, n-1)坐标走右上角上半区，最后1步到达(i,j)收益为0
	// (n-1, 0)坐标走左下角下半区，最后1步到达(i,j)收益为0
	var dp2 = make([]int, n>>1)
	dp2[0] = fruits[0][n-1]
	var dp3 = make([]int, n>>1)
	dp3[0] = fruits[n-1][0]
	
	for i := 1; i < n-1; i++ {
		for k := 0; k <= i; k++ {
			
		}
	}

	var pre3 int
	for j := 1; j < n-1; j++ {
		pre3 = max(dp3[0], dp3[1])
		dp3[0] = pre3 + fruits[n-1][j]
		if j < n-2 {
			dp3[1] = pre3 + fruits[n-2][j]
		}
	}

	return dp1 + dp2[0] + dp3[0]
}
