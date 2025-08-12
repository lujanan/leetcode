// 2438.二的幂数组中查询范围内的乘积

package algorithm_2000

func productQueries(n int, queries [][]int) []int {
	var powers []int
	var num = 1
	for n > 0 {
		if n&1 != 0 {
			powers = append(powers, num)
		}
		n >>= 1
		num <<= 1
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
