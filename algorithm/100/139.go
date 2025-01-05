// 139. 单词拆分

package algorithm_100

func wordBreak(s string, wordDict []string) bool {
	var wordMap = make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}

	var dp = make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[string([]byte(s)[j:i])] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

// dp写得复杂了
func wordBreakV2(s string, wordDict []string) bool {
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
				nextIdx := idx + len(wordDict[i])
				if nextIdx <= len(s) && string(str[idx:nextIdx]) == wordDict[i] {
					if nextIdx == len(s) {
						return true
					}
					if dp[nextIdx] {
						continue
					}
					dp1[nextIdx] = true
				}
			}
		}
		dp = dp1
	}
	return false
}
