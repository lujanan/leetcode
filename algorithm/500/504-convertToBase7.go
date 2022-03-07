//给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
//
//
//
// 示例 1:
//
//
//输入: num = 100
//输出: "202"
//
//
// 示例 2:
//
//
//输入: num = -7
//输出: "-10"
//
//
//
//
// 提示：
//
//
// -10⁷ <= num <= 10⁷
//
// Related Topics 数学 👍 136 👎 0

package algorithm_500

import (
	"math"
	"strconv"
)

func convertToBase7(num int) string {
	var res string
	if num < 0 {
		num = -num
		res = "-"
	}

	var idx, nl = 0, []int{1}
	for {
		idx++
		tn := math.Pow(7, float64(idx))
		if tn > float64(num) {
			break
		}
		nl = append(nl, int(tn))
	}

	for i := len(nl) - 1; i >= 0; i-- {
		n := num / nl[i]
		res += strconv.Itoa(n)
		num -= n * nl[i]
	}

	return res
}
