// 131.分割回文串

package algorithm_100

func partition(s string) [][]string {
	// dp[i][j] = true, s[i-1] = s[j+1], dp[i-1][j+1] = true
	// dp[0][0] = true
	// dp[0][1] = s[0] == s[1]
	// dp[0][2] = s[0] == s[2]
	// dp[0][3] = s[0] == s[3] && s[1] == s[2]
	// dp[0][4] = s[0] == s[4] && s[1] == s[3]
	// dp[0][5] = s[0] == s[5] && s[1] == s[4] && s[2] == s[3]

	// dp[1][2] = s[1] == s[2]
	// dp[1][3] = s[1] == s[3]
	// dp[1][4] = s[1] == s[4] && s[2] == s[3]

	return nil

}