// 72.编辑距离
//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
//
//
// 示例 1：
//
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//
//
// 示例 2：
//
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//
//
//
// 提示：
//
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 3559 👎 0

package algorithm_0

// leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	if len(word1) < 1 {
		return len(word2)
	}
	if len(word2) < 1 {
		return len(word1)
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var dp = make([][]int, len(word1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2))
	}
	if word1[0] != word2[0] {
		dp[0][0] = 1
	}

	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i > 0 && j > 0 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
			} else if i > 0 {
				dp[i][j] = dp[i-1][j]
			} else if j > 0 {
				dp[i][j] = dp[i][j-1]
			}

			if word1[i] != word2[j] {
				dp[i][j]++
			}
		}
	}
	return dp[len(word1)-1][len(word2)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
