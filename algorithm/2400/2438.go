// 2438.二的幂数组中查询范围内的乘积

package algorithm_2000

func productQueries(n int, queries [][]int) []int {
	var maxNum = 1
	for maxNum<<1 <= n {
		maxNum <<= 1
	}

	var powers []int
	for n > 0 {
		if n >= maxNum {
			powers = append(powers, maxNum)
			n -= maxNum
		}
		maxNum >>= 1
	}
	for i, j := 0, len(powers)-1; i < j; i, j = i+1, j-1 {
		powers[i], powers[j] = powers[j], powers[i]
	}

	var ans = make([]int, 0, len(queries))
	var tmp int
	for i := 0; i < len(queries); i++ {
		tmp = 1
		for j := queries[i][0]; j <= queries[i][1] && j < len(powers); j++ {
			tmp = tmp * powers[j] % 1000000007
		}
		ans = append(ans, tmp)
	}

	return ans
}
