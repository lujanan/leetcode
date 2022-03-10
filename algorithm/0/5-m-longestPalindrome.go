package algorithm_0

// 5. 最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring/

// 暴力，已每个字符为中心，向中间两边延伸搜索最长回文串
// O(n^2)
func longestPalindrome(s string) string {
	var n = len(s)
	var fn = func(l, r int) (int, int) {
		for ; l >= 0 && r < n && l <= r && s[l] == s[r]; l, r = l-1, r+1 {
		}
		return l + 1, r - 1
	}

	var begin, ml = 0, 1
	for i := 0; i < n; i++ {
		l, r := fn(i, i)
		if r-l+1 > ml {
			begin, ml = l, r-l+1
		}

		l, r = fn(i, i+1)
		if r-l+1 > ml {
			begin, ml = l, r-l+1
		}
	}
	return string(s[begin : begin+ml])
}

// 动规 O(n^2)
func longestPalindrome1(s string) string {
	var n = len(s)
	if n < 2 {
		return s
	}

	var dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	var begin, ml = 0, 1

	// 子串长度
	for ll := 2; ll <= n; ll++ {
		for l := 0; l < n; l++ {
			var r = l + ll - 1 // 右边界
			if r >= n {
				break
			}

			// if s[l] != s[r] {
			// 	dp[l][r] = false

			// } else {
			// 	if ll < 3 {
			// 		dp[l][r] = true
			// 	} else {
			// 		dp[l][r] = dp[l+1][r-1]
			// 	}
			// }

			// 优化
			dp[l][r] = s[l] == s[r]
			if dp[l][r] && ll >= 3 {
				// 子串是回文串，当前串才能是回文串
				dp[l][r] = dp[l+1][r-1]
			}

			if ml < ll && dp[l][r] {
				begin, ml = l, ll
			}
		}
	}
	return string(s[begin : begin+ml])
}
