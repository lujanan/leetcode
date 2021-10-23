package algorithm_1600

func maximalNetworkRank(n int, roads [][]int) (max int) {
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n+1)
	}
	for i := 0; i < len(roads); i++ {
		res[roads[i][0]][n]++
		res[roads[i][1]][n]++
		res[roads[i][0]][roads[i][1]] = 1
		res[roads[i][1]][roads[i][0]] = 1
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if max < res[i][n]+res[j][n]-res[i][j] {
				max = res[i][n] + res[j][n] - res[i][j]
			}
		}
	}
	return
}
