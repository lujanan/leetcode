// 1688.最大重复子字符串
//给你一个字符串 sequence ，如果字符串 word 连续重复 k 次形成的字符串是 sequence 的一个子字符串，那么单词 word 的 重复值为
// k 。单词 word 的 最大重复值 是单词 word 在 sequence 中最大的重复值。如果 word 不是 sequence 的子串，那么重复值 k
//为 0 。
//
// 给你一个字符串 sequence 和 word ，请你返回 最大重复值 k 。
//
//
//
// 示例 1：
//
//
//输入：sequence = "ababc", word = "ab"
//输出：2
//解释："abab" 是 "ababc" 的子字符串。
//
//
// 示例 2：
//
//
//输入：sequence = "ababc", word = "ba"
//输出：1
//解释："ba" 是 "ababc" 的子字符串，但 "baba" 不是 "ababc" 的子字符串。
//
//
// 示例 3：
//
//
//输入：sequence = "ababc", word = "ac"
//输出：0
//解释："ac" 不是 "ababc" 的子字符串。
//
//
//
//
// 提示：
//
//
// 1 <= sequence.length <= 100
// 1 <= word.length <= 100
// sequence 和 word 都只包含小写英文字母。
//
//
// Related Topics 字符串 动态规划 字符串匹配 👍 167 👎 0

package algorithm_1600

// leetcode submit region begin(Prohibit modification and deletion)
func maxRepeating(sequence string, word string) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dp = make([][3]int, len(sequence))
	for i, j := 0, len(word)-1; i < len(sequence); i, j = i+1, j+1 {
		if j < len(sequence) && sequence[i:j+1] == word {
			dp[i][1], dp[i][2] = 1, 1
		}

		if i > 0 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		}
		if i-len(word) >= 0 && dp[i-len(word)][2] > 0 {
			dp[i][1] += dp[i-len(word)][1]
		}
	}
	return max(dp[len(sequence)-1][0], dp[len(sequence)-1][1])
}

//leetcode submit region end(Prohibit modification and deletion)
