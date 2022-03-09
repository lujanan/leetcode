package algorithm_0

// 5. 最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	var res = string(s[0])
	var dp = make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		dp[i] = 1
		if i-dp[i-1]-1 >= 0 && s[i] == s[i-dp[i-1]-1] {
			dp[i] = dp[i-1] + 2
		} else if i-2 >= 0 && s[i] == s[i-2] {
			dp[i] = 3
		} else if s[i] == s[i-1] {
			dp[i] = 2
		}

		if dp[i] > len(res) {
			res = string(s[i+1-dp[i] : i+1])
		}
	}
	return res
}
