//象棋骑士有一个独特的移动方式，它可以垂直移动两个方格，水平移动一个方格，或者水平移动两个方格，垂直移动一个方格(两者都形成一个 L 的形状)。
//
// 象棋骑士可能的移动方式如下图所示:
//
//
//
// 我们有一个象棋骑士和一个电话垫，如下所示，骑士只能站在一个数字单元格上(即蓝色单元格)。
//
//
//
// 给定一个整数 n，返回我们可以拨多少个长度为 n 的不同电话号码。
//
// 你可以将骑士放置在任何数字单元格上，然后你应该执行 n - 1 次移动来获得长度为 n 的号码。所有的跳跃应该是有效的骑士跳跃。
//
// 因为答案可能很大，所以输出答案模 10⁹ + 7.
//
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 1
//输出：10
//解释：我们需要拨一个长度为1的数字，所以把骑士放在10个单元格中的任何一个数字单元格上都能满足条件。
//
//
// 示例 2：
//
//
//输入：n = 2
//输出：20
//解释：我们可以拨打的所有有效号码为[04, 06, 16, 18, 27, 29, 34, 38, 40, 43, 49, 60, 61, 67, 72,
//76, 81, 83, 92, 94]
//
//
// 示例 3：
//
//
//输入：n = 3131
//输出：136006598
//解释：注意取模
//
//
//
//
// 提示：
//
//
// 1 <= n <= 5000
//
//
// Related Topics 动态规划 👍 177 👎 0

package algorithm_900

// leetcode submit region begin(Prohibit modification and deletion)
func knightDialer(n int) int {
	var num = []int{2, 2, 2, 2, 3, 0, 3, 2, 2, 2}
	var jump = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for j := 1; j < n; j++ {
		for i := 0; i < 10; i++ {
			jump[i] = jump[i] * num[i]
		}
	}

	for i := 1; i < 10; i++ {
		jump[0] += jump[i]
	}
	return jump[0] % (1000000000 + 7)
}

//leetcode submit region end(Prohibit modification and deletion)
