// 624.数组列表中的最大距离

package algorithm_600

import "math"

func maxDistance(arrays [][]int) int {
	var res int
	var dp = [2]int{arrays[0][0], arrays[0][len(arrays[0])-1]}
	var abs = func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	for i := 1; i < len(arrays); i++ {
		res = max(res, abs(arrays[i][0], dp[1]), abs(arrays[i][len(arrays[i])-1], dp[0]))
		dp[0] = min(dp[0], arrays[i][0])
		dp[1] = max(dp[1], arrays[i][len(arrays[i])-1])
	}

	return res
}

func maxDistance1(arrays [][]int) int {
	var dp = [4]int{0, 1, 0, 1}
	if arrays[1][0] < arrays[0][0] {
		dp[0], dp[1] = dp[1], dp[0]
	}
	if arrays[1][len(arrays[1])-1] > arrays[0][len(arrays[0])-1] {
		dp[2], dp[3] = dp[3], dp[2]
	}

	for i := 2; i < len(arrays); i++ {
		if arrays[i][0] < arrays[dp[0]][0] {
			dp[0], dp[1] = i, dp[0]
		} else if arrays[i][0] < arrays[dp[1]][0] {
			dp[1] = i
		}
		if arrays[i][len(arrays[i])-1] > arrays[dp[2]][len(arrays[dp[2]])-1] {
			dp[2], dp[3] = i, dp[2]
		} else if arrays[i][len(arrays[i])-1] > arrays[dp[3]][len(arrays[dp[3]])-1] {
			dp[3] = i
		}
	}

	isSame := dp[0] == dp[2]
	dp[0] = arrays[dp[0]][0]
	dp[1] = arrays[dp[1]][0]
	dp[2] = arrays[dp[2]][len(arrays[dp[2]])-1]
	dp[3] = arrays[dp[3]][len(arrays[dp[3]])-1]

	if !isSame {
		return int(math.Abs(float64(dp[2] - dp[0])))
	}
	return int(max(math.Abs(float64(dp[3]-dp[0])), math.Abs(float64(dp[2]-dp[1]))))
}
