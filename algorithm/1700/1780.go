// 1780.判断一个数字是否可以表示成三的幂的和

package algorithm_1700

import "math"

// 转为3进制计算
func checkPowersOfThree(n int) bool {
	// 求余=2则无法表示为3的n次幂
	// 求余=1则可表示为3的0次幂
	for ; n > 0; n /= 3 {
		if n%3 == 2 {
			return false
		}
	}
	return true
}

// 当 m = 3^z 时，必有 m = 3^x * 3^y, x+y=z
// n<=10^7, 所以 z 的最大值为14
// n = n - 3^z - 3^(z-1) ... z^0
// return n == 0
func checkPowersOfThree0(n int) bool {
	var base = 14
	var num int
	for n > 0 && base >= 0 {
		num = int(math.Pow(3, float64(base)))
		if n >= num {
			n -= num
		}
		base--
	}
	return n == 0
}
