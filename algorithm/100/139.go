// 139. 单词拆分

package algorithm_100

func wordBreak(s string, wordDict []string) bool {
	var str = []byte(s)
	var dp = make(map[int]bool)
	for i := 0; i < len(wordDict); i++ {
		if s == wordDict[i] {
			return true
		}

		if len(wordDict[i]) > len(s) {
			continue
		}

		if string(str[:len(wordDict[i])]) == wordDict[i] {
			dp[len(wordDict[i])] = true
		}
	}

	for len(dp) > 0 {
		var dp1 = make(map[int]bool)
		for i := 0; i < len(wordDict); i++ {
			for idx := range dp {
				if idx+len(wordDict[i]) <= len(s) &&
					string(str[idx:idx+len(wordDict[i])]) == wordDict[i] {
					if idx+len(wordDict[i]) == len(s) {
						return true
					}
					dp1[idx+len(wordDict[i])] = true
				}
			}
		}
		dp = dp1
	}
	return false
}
