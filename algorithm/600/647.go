// 647.回文子串
//给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//
// 回文字符串 是正着读和倒过来读一样的字符串。
//
// 子字符串 是字符串中的由连续字符组成的一个序列。
//
//
//
// 示例 1：
//
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
//
// Related Topics 双指针 字符串 动态规划 👍 1403 👎 0

package algorithm_600

// leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	var dp = make([]int, len(s))
	var res = 1
	for i := 1; i < len(s); i++ {
		res += 1
		if s[i] == s[i-1] {
			res++
			dp[i] = 2
		
		}
		if i > 1 && s[i] == s[i-2] {
			res++
		}
		if dp[i-1] > 0 && i-dp[i-1] >= 0 && s[i-dp[i-1]] == s[i] {
			res++
			dp[i] = dp[i-1] + 2
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
