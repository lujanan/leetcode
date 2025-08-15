// 342.4的幂

package algorithm_300

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
