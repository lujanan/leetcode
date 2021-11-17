//给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为
//每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
//
//
//
// 示例 1:
//
//
//输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
//输出: 16
//解释: 这两个单词为 "abcw", "xtfn"。
//
// 示例 2:
//
//
//输入: ["a","ab","abc","d","cd","bcd","abcd"]
//输出: 4
//解释: 这两个单词为 "ab", "cd"。
//
// 示例 3:
//
//
//输入: ["a","aa","aaa","aaaa"]
//输出: 0
//解释: 不存在这样的两个单词。
//
//
//
//
// 提示：
//
//
// 2 <= words.length <= 1000
// 1 <= words[i].length <= 1000
// words[i] 仅包含小写字母
//
// Related Topics 位运算 数组 字符串 👍 259 👎 0

package algorithm_300

func maxProduct(words []string) (num int) {
	var wordMap = make(map[int]int)
	for i := 0; i < len(words); i++ {
		var tmp = 0
		for j := 0; j < len(words[i]); j++ {
			tmp = tmp | 1<<int((words[i][j]-'a'))
		}
		if n, ok := wordMap[tmp]; !ok || n < len(words[i]) {
			wordMap[tmp] = len(words[i])
		}
	}

	for k, n := range wordMap {
		for k1, n1 := range wordMap {
			if k1&k == 0 && n*n1 > num {
				num = n * n1
			}
		}
	}
	return
}
