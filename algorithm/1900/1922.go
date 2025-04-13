// 1922.统计好数字的数目

package algorithm_1900

import "math"

func countGoodNumbers(n int64) int {
	var mod = 1000000007
	num4, num5 := (n+1)/2, n/2

	res := int64(math.Pow(float64(4), float64(num4))) % int64(mod)
	res = int64(math.Pow(float64(5), float64(num5))) % int64(mod)

	// for i := int64(0); i < n; i++ {
	// 	if i&1 != 0 {
	// 		res = res * 4 % mod
	// 	} else {
	// 		res = res * 5 % mod
	// 	}
	// }
	return int(res)
}
