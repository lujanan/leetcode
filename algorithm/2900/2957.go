// 2957.消除相邻近似相等字符
//给你一个下标从 0 开始的字符串 word 。
//
// 一次操作中，你可以选择 word 中任意一个下标 i ，将 word[i] 修改成任意一个小写英文字母。
//
// 请你返回消除 word 中所有相邻 近似相等 字符的 最少 操作次数。
//
// 两个字符 a 和 b 如果满足 a == b 或者 a 和 b 在字母表中是相邻的，那么我们称它们是 近似相等 字符。
//
//
//
// 示例 1：
//
//
//输入：word = "aaaaa"
//输出：2
//解释：我们将 word 变为 "acaca" ，该字符串没有相邻近似相等字符。
//消除 word 中所有相邻近似相等字符最少需要 2 次操作。
//
//
// 示例 2：
//
//
//输入：word = "abddez"
//输出：2
//解释：我们将 word 变为 "ybdoez" ，该字符串没有相邻近似相等字符。
//消除 word 中所有相邻近似相等字符最少需要 2 次操作。
//
// 示例 3：
//
//
//输入：word = "zyxyxyz"
//输出：3
//解释：我们将 word 变为 "zaxaxaz" ，该字符串没有相邻近似相等字符。
//消除 word 中所有相邻近似相等字符最少需要 3 次操作
//
//
//
//
// 提示：
//
//
// 1 <= word.length <= 100
// word 只包含小写英文字母。
//
//
// Related Topics 贪心 字符串 动态规划 👍 13 👎 0

package algorithm_2900

// leetcode submit region begin(Prohibit modification and deletion)
func removeAlmostEqualCharacters(word string) int {
	// word[i] == word[i-1]  dp[i][0] = dp[i-1][1]    dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
	// word[i] != word[i-1]  dp[i][0] = min(dp[i-1][0], dp[i-1][1])    dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1

	var diff = func(a, b byte) int {
		if a > b {
			return int(a - b)
		} else {
			return int(b - a)
		}
	}

	var dp = make([]int, 2)
	dp[1] = 1
	for i := 1; i < len(word); i++ {
		if diff(word[i], word[i-1]) <= 1 {
			// 相邻字符近似相等，选择变当前字符，或前1字符
			// 1.当前字符不变，则前1字符必须变
			// 2.当前字符变，则前1字符变或不变都可，取最小操作数 + 1(当前操作数)
			dp[0], dp[1] = dp[1], min(dp[0], dp[1])+1
		} else {
			// 相邻字符不是近似相等，不用变，取最小操作数即可
			dp[0], dp[1] = min(dp[0], dp[1]), min(dp[0], dp[1])+1
		}
	}
	return min(dp[0], dp[1])
}

//leetcode submit region end(Prohibit modification and deletion)
