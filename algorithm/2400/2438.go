// 2438.二的幂数组中查询范围内的乘积

package algorithm_2000

import "math"

func productQueries(n int, queries [][]int) []int {
	var maxNum = 1
	var twoNum int
	for maxNum<<1 <= n {
		maxNum <<= 1
		twoNum++
	}

	var powers []int
	for n > 0 {
		if n >= maxNum {
			powers = append(powers, twoNum)
			n -= maxNum
		}
		maxNum >>= 1
		twoNum--
	}

	var p = make([]int, len(powers))
	for i, j := len(powers)-1, 0; i >= 0; i, j = i-1, j+1 {
		p[j] = powers[i]
		if j > 0 {
			p[j] += p[j-1]
		}
	}

	var ans = make([]int, 0, len(queries))
	var tmp int
	for i := 0; i < len(queries); i++ {
		tmp = p[queries[i][1]]
		if 0 < queries[i][0] {
			tmp -= p[queries[i][0]-1]
		}

		var num = math.Pow(2, float64(tmp))
		ans = append(ans, int(math.Mod(num, 1000000007)))
	}

	return ans
}
