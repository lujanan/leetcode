//给出一个字符串数组 words 组成的一本英语词典。返回 words 中最长的一个单词，该单词是由 words 词典中其他单词逐步添加一个字母组成。
//
// 若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。
//
//
//
// 示例 1：
//
//
//输入：words = ["w","wo","wor","worl", "world"]
//输出："world"
//解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。
//
//
// 示例 2：
//
//
//输入：words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
//输出："apple"
//解释："apply" 和 "apple" 都能由词典中的单词组成。但是 "apple" 的字典序小于 "apply"
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 1000
// 1 <= words[i].length <= 30
// 所有输入的字符串 words[i] 都只包含小写字母。
//
// Related Topics 字典树 数组 哈希表 字符串 排序 👍 265 👎 0

package algorithm_700

import "sort"

// todo 字典树解法

// 排序+hash 优化
func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return true
		} else if len(words[i]) == len(words[j]) {
			return words[i] > words[j]
		}
		return false
	})
	var wMap = make(map[string]int)
	var longWord string
	for i := range words {
		if len(words[i]) == 1 {
			longWord = words[i]
			wMap[words[i]] = 0
		} else if len(words[i]) > 1 {
			if _, ok := wMap[words[i][:len(words[i])-1]]; ok {
				longWord = words[i]
				wMap[words[i]] = 0
			}
		}
	}
	return longWord
}

// 排序+hash
func longestWord1(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) > len(words[j]) {
			return true
		}
		if len(words[i]) == len(words[j]) {
			for k := 0; k < len(words[i]); k++ {
				if words[i][k] == words[j][k] {
					continue
				} else if words[i][k] < words[j][k] {
					return true
				}
				return false
			}
			return i < j
		}
		return false
	})

	var wMap = make(map[string]int)
	for i := range words {
		wMap[words[i]] = 0
	}

	for i := range words {
		var found bool
		for j := len(words[i]) - 1; j >= 0; j-- {
			if j == 0 {
				found = true
				break
			}
			if _, ok := wMap[words[i][:j]]; !ok {
				break
			}
		}

		if found {
			return words[i]
		}
	}

	return ""
}
