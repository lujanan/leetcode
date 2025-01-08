// 97.交错字符串
//给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
//
// 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//
//
// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
//
//
// 注意：a + b 意味着字符串 a 和 b 连接。
//
//
//
// 示例 1：
//
//
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出：true
//
//
// 示例 2：
//
//
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出：false
//
//
// 示例 3：
//
//
//输入：s1 = "", s2 = "", s3 = ""
//输出：true
//
//
//
//
// 提示：
//
//
// 0 <= s1.length, s2.length <= 100
// 0 <= s3.length <= 200
// s1、s2、和 s3 都由小写英文字母组成
//
//
//
//
// 进阶：您能否仅使用 O(s2.length) 额外的内存空间来解决它?
//
// Related Topics 字符串 动态规划 👍 1063 👎 0

package algorithm_0

// leetcode submit region begin(Prohibit modification and deletion)
func isInterleave(s1 string, s2 string, s3 string) bool {

	// dp状态
	// 如果 k = i + j - 1, s3[:k] 由 s1[:i] 和 s2[:j] 交错构成
	// 那必然存在 s3[k-1] == s1[i-1] 或 s3[k-1] == s2[j-1]，即 s3 的最后一个字符必定等于 si 或 s2 的最后一个字符 
	// 所以 s3[k-1] = s3[k-2] + s1[i-1] || s3[k-2] + s2[j-1]
	// 只要知道 s3[k-2] 的状态，即可根据 s1 或 s2 判断 s3[k-1] 的状态
	// s3[0] = s1[0] || s2[0]
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	var dp = make([]bool, len(s2)+1)
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[j] = true
				continue
			}

			dp[j] = (i > 0 && dp[j] && s1[i-1] == s3[i+j-1]) || (j > 0 && dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[len(s2)]
}

func isInterleaveV3(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	var dp = make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

func isInterleaveV2(s1 string, s2 string, s3 string) bool {
	var dpMap = make(map[[2]int]struct{})
	dpMap[[2]int{-1, -1}] = struct{}{}

	for i := 0; i < len(s3); i++ {
		var tmpMap = make(map[[2]int]struct{})
		for dp := range dpMap {
			if dp[0]+1 < len(s1) && s1[dp[0]+1] == s3[i] {
				tmpMap[[2]int{dp[0] + 1, dp[1]}] = struct{}{}
			}
			if dp[1]+1 < len(s2) && s2[dp[1]+1] == s3[i] {
				tmpMap[[2]int{dp[0], dp[1] + 1}] = struct{}{}
			}
		}
		dpMap = tmpMap
	}

	for dp := range dpMap {
		if dp[0]+dp[1]+2 == len(s1)+len(s2) {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
