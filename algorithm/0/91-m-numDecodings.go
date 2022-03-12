package algorithm_0

// 91. 解码方法
// https://leetcode-cn.com/problems/decode-ways/

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	var dp0, dp1, dp2 = 1, 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			dp2 = 0
		}
		
		if int(s[i-1]-'0')*10+int(s[i]-'0') >= 10 &&
			int(s[i-1]-'0')*10+int(s[i]-'0') <= 26 {
			dp2 += dp0
		}
		dp0, dp1 = dp1, dp2
	}
	return dp2
}
