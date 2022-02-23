//给你一个字符串 s ，根据下述规则反转字符串：
//
//
// 所有非英文字母保留在原有位置。
// 所有英文字母（小写或大写）位置反转。
//
//
// 返回反转后的 s 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "ab-cd"
//输出："dc-ba"
//
//
//
//
//
// 示例 2：
//
//
//输入：s = "a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba"
//
//
//
//
//
// 示例 3：
//
//
//输入：s = "Test1ng-Leet=code-Q!"
//输出："Qedo1ct-eeLg=ntse-T!"
//
//
//
//
// 提示
//
//
// 1 <= s.length <= 100
// s 仅由 ASCII 值在范围 [33, 122] 的字符组成
// s 不含 '\"' 或 '\\'
//
// Related Topics 双指针 字符串 👍 145 👎 0

package algorithm_900

func reverseOnlyLetters(s string) string {
	var str = []byte(s)
	var l, r = 0, len(str) - 1
	for l < r {
		if str[l] < 'A' || ('Z' < str[l] && str[l] < 'a') || 'z' < str[l] {
			l++
			continue
		}
		if str[r] < 'A' || ('Z' < str[r] && str[r] < 'a') || 'z' < str[r] {
			r--
			continue
		}
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
	return string(str)
}
