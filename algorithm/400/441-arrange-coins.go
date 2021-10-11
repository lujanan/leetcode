//你总共有 n 枚硬币，并计划将它们按阶梯状排列。对于一个由 k 行组成的阶梯，其第 i 行必须正好有 i 枚硬币。阶梯的最后一行 可能 是不完整的。
//
// 给你一个数字 n ，计算并返回可形成 完整阶梯行 的总行数。
//
//
//
// 示例 1：
//
//
//输入：n = 5
//输出：2
//解释：因为第三行不完整，所以返回 2 。
//
//
// 示例 2：
//
//
//输入：n = 8
//输出：3
//解释：因为第四行不完整，所以返回 3 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
// Related Topics 数学 二分查找 👍 174 👎 0

package algorithm_400

func arrangeCoins(n int) int {
	var l, r = 1, n
	for r >= l {
		var m = l + (r-l)>>1
		var mt = (1 + m) * (m >> 1)
		if m%2 == 1 {
			mt += m>>1 + 1
		}
		if mt == n || (mt < n && mt+m+1 > n) {
			return m
		} else if mt > n && mt-m <= n {
			return m - 1
		}

		if mt > n {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return 1
}
