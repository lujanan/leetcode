//统计所有小于非负整数 n 的质数的数量。
//
//
//
// 示例 1：
//
// 输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
//
//
// 示例 2：
//
// 输入：n = 0
//输出：0
//
//
// 示例 3：
//
// 输入：n = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= n <= 5 * 106
//
// Related Topics 数组 数学 枚举 数论
// 👍 736 👎 0

package algorithm

func countPrimes(n int) int {
	var arr = make([]int, n+1)
	for i := 2; i < n; i++ {
		arr[i] = 1
	}
	var count = 0
	for i := 2; i < n; i++ {
		if arr[i] == 1 {
			count++
			for j := i * i; j < n; j = j + i {
				arr[j] = 0
			}
		}
	}
	return count
}

// 超时警告
func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes1(n int) (cnt int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return
}
