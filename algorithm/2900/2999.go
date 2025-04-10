// 2999.统计强大整数的数目
//给你三个整数 start ，finish 和 limit 。同时给你一个下标从 0 开始的字符串 s ，表示一个 正 整数。
//
// 如果一个 正 整数 x 末尾部分是 s （换句话说，s 是 x 的 后缀），且 x 中的每个数位至多是 limit ，那么我们称 x 是 强大的 。
//
// 请你返回区间 [start..finish] 内强大整数的 总数目 。
//
// 如果一个字符串 x 是 y 中某个下标开始（包括 0 ），到下标为 y.length - 1 结束的子字符串，那么我们称 x 是 y 的一个后缀。比方说，
//25 是 5125 的一个后缀，但不是 512 的后缀。
//
//
//
// 示例 1：
//
//
//输入：start = 1, finish = 6000, limit = 4, s = "124"
//输出：5
//解释：区间 [1..6000] 内的强大数字为 124 ，1124 ，2124 ，3124 和 4124 。这些整数的各个数位都 <= 4 且 "124"
//是它们的后缀。注意 5124 不是强大整数，因为第一个数位 5 大于 4 。
//这个区间内总共只有这 5 个强大整数。
//
//
// 示例 2：
//
//
//输入：start = 15, finish = 215, limit = 6, s = "10"
//输出：2
//解释：区间 [15..215] 内的强大整数为 110 和 210 。这些整数的各个数位都 <= 6 且 "10" 是它们的后缀。
//这个区间总共只有这 2 个强大整数。
//
//
// 示例 3：
//
//
//输入：start = 1000, finish = 2000, limit = 4, s = "3000"
//输出：0
//解释：区间 [1000..2000] 内的整数都小于 3000 ，所以 "3000" 不可能是这个区间内任何整数的后缀。
//
//
//
//
// 提示：
//
//
// 1 <= start <= finish <= 10¹⁵
// 1 <= limit <= 9
// 1 <= s.length <= floor(log10(finish)) + 1
// s 数位中每个数字都小于等于 limit 。
// s 不包含任何前导 0 。
//
//
// Related Topics 数学 字符串 动态规划 👍 49 👎 0

package algorithm_2900

import (
	"math"
	"strconv"
)

// leetcode submit region begin(Prohibit modification and deletion)

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
    start_ := strconv.FormatInt(start - 1, 10)
	finish_ := strconv.FormatInt(finish, 10)
	return calculate(finish_, s, limit) - calculate(start_, s, limit)
}

func calculate(x string, s string, limit int) int64 {
	if len(x) < len(s) {
		return 0
	}
	if len(x) == len(s) {
		if x >= s {
			return 1
		}
		return 0
	}

	suffix := x[len(x) - len(s):]
	var count int64
	preLen := len(x) - len(s)

	for i := 0; i < preLen; i++ {
		digit := int(x[i] - '0')
		if limit < digit {
			count += int64(math.Pow(float64(limit + 1), float64(preLen - i)))
			return count
		}
		count += int64(digit) * int64(math.Pow(float64(limit + 1), float64(preLen - 1- i)))
	}
	if suffix >= s {
		count++
	}
	return count
}


/**
func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {

	var l = len(s)
	var num, _ = strconv.ParseInt(s, 10, 64)

	var fl int
	var fArr = make([]int64, 0)
	var tmpF = finish
	for tmpF > 0 {
		fl++
		fArr = append(fArr, tmpF%10)
		tmpF /= 10
	}
	var sl int
	var tmpS = start
	for tmpS > 0 {
		sl++
		tmpS /= 10
	}

	// 有pre前缀
	var pre = fl - max(l, sl)
	if l >= fl || pre <= 0 {
		if start <= num && num <= finish {
			return 1
		}
		return 0
	}

	var res int64
	if num >= start {
		res++
	}

	// 计算第[1, pre-1]前缀的合法数量
	for i := 1; i < pre; i++ {
		res += int64(limit) * int64(math.Pow(float64(limit+1), float64(i-1)))
	}

	// 计算第pre前缀的合法数量
	base := min(int64(limit), fArr[fl-1]) 
	if int64(limit) >= fArr[fl-1] {
		base -= 1
		res += base * int64(math.Pow(float64(limit+1), float64(pre-1)))

		for i := 2; i <= pre; i++ {
			base1 := min(int64(limit), fArr[fl-i]) + 1
			res += base1 * int64(math.Pow(float64(limit+1), float64(pre-i)))
		}
	} else {
		res += base * int64(math.Pow(float64(limit+1), float64(pre-1)))
	}
	return res
}
*/

//leetcode submit region end(Prohibit modification and deletion)
