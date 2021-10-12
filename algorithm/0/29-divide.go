//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
// 返回被除数 dividend 除以除数 divisor 得到的商。
//
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//
//
//
// 示例 1:
//
// 输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
//
// 示例 2:
//
// 输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = truncate(-2.33333..) = -2
//
//
//
// 提示：
//
//
// 被除数和除数均为 32 位有符号整数。
// 除数不为 0。
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2³¹, 231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
//
// Related Topics 位运算 数学 👍 714 👎 0

package algorithm_0

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	} else if dividend == divisor {
		return 1
	}

	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		} else if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 {
		return 0
	}
	// 统一转化为负数计算
	var rev = false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	var l, r, res = 1, math.MaxInt32, 0
	for l <= r {
		mid := l + (r-l)>>1
		check := quickAdd(dividend, divisor, mid)
		if check {
			res = mid
			if mid == math.MaxInt32 {
				break
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if rev {
		res = -res
	}
	return res
}

// 快速加
func quickAdd(dividend, divisor, mid int) bool {
	var res, add = 0, divisor
	for mid > 0 {
		if mid&1 != 0 {
			// 比被除数大，不符合
			if res+add < dividend {
				return false
			}
			res += add
			mid--
		} else {
			// 比被除数大，不符合
			if add+add < dividend {
				return false
			}
			add += add
			mid >>= 1
		}
	}
	return true
}
