//复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：
//
//
// 实部 是一个整数，取值范围是 [-100, 100]
// 虚部 也是一个整数，取值范围是 [-100, 100]
// i² == -1
//
//
// 给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。
//
//
//
// 示例 1：
//
//
//输入：num1 = "1+1i", num2 = "1+1i"
//输出："0+2i"
//解释：(1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
//
//
// 示例 2：
//
//
//输入：num1 = "1+-1i", num2 = "1+-1i"
//输出："0+-2i"
//解释：(1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。
//
//
//
//
// 提示：
//
//
// num1 和 num2 都是有效的复数表示。
//
// Related Topics 数学 字符串 模拟 👍 101 👎 0

package algorithm_500

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	var fn = func(str string) (int, int) {
		str = strings.Replace(str, "i", "", -1)
		if strArr := strings.Split(str, "+"); len(strArr) >= 2 {
			n1, _ := strconv.Atoi(strArr[len(strArr)-2])
			n2, _ := strconv.Atoi(strArr[len(strArr)-1])
			return n1, n2
		}
		return 0, 0
	}

	n1, n2 := fn(num1)
	m1, m2 := fn(num2)
	return strconv.Itoa(n1*m1-n2*m2) + "+" + strconv.Itoa(n1*m2+n2*m1) + "i"
}
