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
	const mod = 1000000000 + 7
	var num = [2][10]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
	for i := 2; i <= n; i++ {
		// 数值可能很大，取模
		num[1][0] = (num[0][4] + num[0][6]) % mod
		num[1][1] = (num[0][6] + num[0][8]) % mod
		num[1][2] = (num[0][7] + num[0][9]) % mod
		num[1][3] = (num[0][4] + num[0][8]) % mod
		num[1][4] = (num[0][0] + num[0][3] + num[0][9]) % mod
		num[1][5] = 0
		num[1][6] = (num[0][0] + num[0][1] + num[0][7]) % mod
		num[1][7] = (num[0][2] + num[0][6]) % mod
		num[1][8] = (num[0][1] + num[0][3]) % mod
		num[1][9] = (num[0][2] + num[0][4]) % mod

		num[0] = num[1]
	}

	for i := 1; i < 10; i++ {
		num[0][0] += num[0][i]
	}
	return num[0][0] % mod
}

//leetcode submit region end(Prohibit modification and deletion)
