//给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
//
// 注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-
//distinct-characters 相同
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
// 1 <= s.length <= 10⁴
// s 由小写英文字母组成
//
// Related Topics 栈 贪心 字符串 单调栈 👍 633 👎 0

package algorithm_300

func removeDuplicateLetters(s string) string {
	var lmap = [26]rune{}
	for _, v := range s {
		lmap[v-'a']++
	}

	var inStack = [26]bool{}
	var str []byte
	for i := 0; i < len(s); i++ {
		var ch = s[i]
		if !inStack[ch-'a'] {
			for len(str) > 0 && ch < str[len(str)-1] {
				last := str[len(str)-1] - 'a'
				if lmap[last] < 1 {
					break
				}
				str = str[:len(str)-1]
				inStack[last] = false
			}
			str = append(str, ch)
			inStack[ch-'a'] = true
		}
		lmap[ch-'a']--
	}
	return string(str)
}
