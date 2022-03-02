//给定一个表示整数的字符串 n ，返回与它最近的回文整数（不包括自身）。如果不止一个，返回较小的那个。
//
// “最近的”定义为两个整数差的绝对值最小。
//
//
//
// 示例 1:
//
//
//输入: n = "123"
//输出: "121"
//
//
// 示例 2:
//
//
//输入: n = "1"
//输出: "0"
//解释: 0 和 2是最近的回文，但我们返回最小的，也就是 0。
//
//
//
//
// 提示:
//
//
// 1 <= n.length <= 18
// n 只由数字组成
// n 不含前导 0
// n 代表在 [1, 10¹⁸ - 1] 范围内的整数
//
// Related Topics 数学 字符串 👍 186 👎 0

package algorithm_500

import (
	"math"
	"strconv"
)

// 直接模拟
func nearestPalindromic(n string) string {
	var l = len(n)
	var list = []int{int(math.Pow10(l-1) - 1), int(math.Pow10(l) + 1)}
	prefix, _ := strconv.Atoi(n[:(l+1)/2])
	pres := []int{prefix - 1, prefix, prefix + 1}

	for _, v := range pres {
		var x = v
		if l&1 == 1 {
			x = x / 10
		}
		for ; x > 0; x /= 10 {
			v = v*10 + x%10
		}
		list = append(list, v)
	}

	var abs = func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}
	var res = -1
	num, _ := strconv.Atoi(n)
	for _, v := range list {
		if num == v {
			continue
		}
		if res == -1 ||
			abs(num-res) > abs(num-v) ||
			(abs(num-res) == abs(num-v) && res > v) {
			res = v
		}
	}
	return strconv.Itoa(res)
}
