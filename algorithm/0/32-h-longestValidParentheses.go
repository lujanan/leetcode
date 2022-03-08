//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
//
// Related Topics 栈 字符串 动态规划 👍 1685 👎 0

package algorithm_0

func longestValidParentheses(s string) int {
	var res, left int
	var dp = make([]int, len(s)+1)
	for i := range s {
		if s[i] == '(' {
			left++
			continue
		}

		if left > 0 {
			dp[i+1] = dp[i] + 2
			dp[i+1] += dp[i+1-dp[i+1]]
			left--
			if res < dp[i+1] {
				res = dp[i+1]
			}
		}
	}

	return res
}
