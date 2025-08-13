// 326.3的幂

package algorithm_300

func isPowerOfThree(n int) bool {
	for n >= 3 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return n == 1
}
