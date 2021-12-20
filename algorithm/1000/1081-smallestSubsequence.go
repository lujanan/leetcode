//返回 s 字典序最小的子序列，该子序列包含 s 的所有不同字符，且只包含一次。
//
// 注意：该题与 316 https://leetcode.com/problems/remove-duplicate-letters/ 相同
//
//
//
// 示例 1：
//
//
//输入：s = "bcabc"
//输出："abc"
//
//
// 示例 2：
//
//
//输入：s = "cbacdcbc"
//输出："acdb"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
// Related Topics 栈 贪心 字符串 单调栈 👍 131 👎 0

package algorithm_1000

func smallestSubsequence(s string) string {
	var lmap = [26]int{}
	for _, v := range s {
		lmap[v-'a']++
	}

	var stack = [26]bool{}
	var str []byte
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if !stack[ch-'a'] {
			for len(str) > 0 && ch < str[len(str)-1] {
				last := str[len(str)-1] - 'a'
				if lmap[last] < 1 {
					break
				}
				str = str[:len(str)-1]
				stack[last] = false
			}

			str = append(str, ch)
			stack[ch-'a'] = true
		}
		lmap[ch-'a']--
	}
	return string(str)
}
