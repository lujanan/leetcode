//给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
//
// 如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
//输出："apple"
//
//
// 示例 2：
//
//
//输入：s = "abpcplea", dictionary = ["a","b","c"]
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// 1 <= dictionary.length <= 1000
// 1 <= dictionary[i].length <= 1000
// s 和 dictionary[i] 仅由小写英文字母组成
//
// Related Topics 数组 双指针 字符串 排序
// 👍 231 👎 0

package algorithm_500

import (
	"strings"
)

func findLongestWord(s string, dictionary []string) (res string) {
	for i := 0; i < len(dictionary); i++ {
		if len(res) > len(dictionary[i]) {
			continue
		}
		m := len(dictionary[i]) - 1
		for n := len(s) - 1; n >= 0 && m >= 0; {
			if s[n] == dictionary[i][m] {
				m--
			}
			n--
		}
		if m == -1 &&
			(res == "" ||
				len(dictionary[i]) > len(res) ||
				(len(dictionary[i]) == len(res) && strings.Compare(res, dictionary[i]) > 0)) {
			res = dictionary[i]
		}
	}
	return
}
