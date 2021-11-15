//初始时有 n 个灯泡处于关闭状态。第一轮，你将会打开所有灯泡。接下来的第二轮，你将会每两个灯泡关闭一个。
//
// 第三轮，你每三个灯泡就切换一个灯泡的开关（即，打开变关闭，关闭变打开）。第 i 轮，你每 i 个灯泡就切换一个灯泡的开关。直到第 n 轮，你只需要切换最后
//一个灯泡的开关。
//
// 找出并返回 n 轮后有多少个亮着的灯泡。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 3
//输出：1
//解释：
//初始时, 灯泡状态 [关闭, 关闭, 关闭].
//第一轮后, 灯泡状态 [开启, 开启, 开启].
//第二轮后, 灯泡状态 [开启, 关闭, 开启].
//第三轮后, 灯泡状态 [开启, 关闭, 关闭].
//
//你应该返回 1，因为只有一个灯泡还亮着。
//
//
// 示例 2：
//
//
//输入：n = 0
//输出：0
//
//
// 示例 3：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁹
//
// Related Topics 脑筋急转弯 数学 👍 263 👎 0

package algorithm_300

import "math"

// 在第 i 轮时，我们会将所有编号为 i 的倍数的灯泡进行切换
// 对于第 k 个灯泡，它被切换的次数恰好就是 k 的约数个数。
// 如果 k 有偶数个约数，那么最终第 k 个灯泡的状态为暗；如果 k 有奇数个约数，那么最终第 k 个灯泡的状态为亮
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}