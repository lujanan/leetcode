// 2209.用地毯覆盖后的最少白色砖块

package algorithm_2200

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	var res = make([][]int, 2)
	res[0] = make([]int, len(floor)+1)

	for i := 0; i < len(floor); i++ {
		res[0][i+1] = res[0][i]
		if floor[i] == '1' {
			res[0][i+1] += 1
		}
	}

	return res[0][len(floor)]
}

// i > 0 时, dp[i][j] = min(dp[i][j-1] + f[j], dp[i-1][j-carpetLen])
// i = 0 时, dp[i][j] = dp[i][j-1] + f[j]
func minimumWhiteTilesV2(floor string, numCarpets int, carpetLen int) int {
	var f = make([]int, len(floor)+1)
	for k, v := range floor {
		if v == '1' {
			f[k+1] = 1
		}
	}

	var exists = make([][]int, numCarpets+1)
	for i := 0; i < len(exists); i++ {
		exists[i] = make([]int, len(floor)+1)
		for j := 0; j < len(floor)+1; j++ {
			exists[i][j] = -1
		}
	}

	var dfs func(num, j int) int
	dfs = func(num, j int) int {
		// 没有格子 或 只要剩下一条毯就能覆盖全部
		if j <= 0 || (num > 0 && j-carpetLen <= 0) {
			return 0
		}

		// 已计算过
		if exists[num][j] >= 0 {
			return exists[num][j]
		}

		exists[num][j] = dfs(num, j-1) + f[j]
		if num > 0 {
			exists[num][j] = min(exists[num][j], dfs(num-1, j-carpetLen))
		}
		return exists[num][j]
	}

	return dfs(numCarpets, len(floor))
}
