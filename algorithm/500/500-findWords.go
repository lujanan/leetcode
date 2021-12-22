//给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。
//
// 美式键盘 中：
//
//
// 第一行由字符 "qwertyuiop" 组成。
// 第二行由字符 "asdfghjkl" 组成。
// 第三行由字符 "zxcvbnm" 组成。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：words = ["Hello","Alaska","Dad","Peace"]
//输出：["Alaska","Dad"]
//
//
// 示例 2：
//
//
//输入：words = ["omk"]
//输出：[]
//
//
// 示例 3：
//
//
//输入：words = ["adsdf","sfd"]
//输出：["adsdf","sfd"]
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 20
// 1 <= words[i].length <= 100
// words[i] 由英文字母（小写和大写字母）组成
//
// Related Topics 数组 哈希表 字符串 👍 193 👎 0

package algorithm_500

import "strings"

func findWords(words []string) []string {
	var wMap = [26]int{2, 3, 3, 2, 1, 2, 2, 2, 1, 2, 2, 2, 3, 3, 1, 1, 1, 1, 2, 1, 1, 3, 1, 3, 1, 3}
	var res []string
	for _, word := range words {
		tmpWord := strings.ToLower(word)
		idx := wMap[tmpWord[0]-'a']
		for _, ch := range tmpWord {
			if wMap[ch-'a'] != idx {
				idx = wMap[ch-'a']
				break
			}
		}
		if idx == wMap[tmpWord[0]-'a'] {
			res = append(res, word)
		}
	}
	return res
}
