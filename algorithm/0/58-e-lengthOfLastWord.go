package algorithm_0

// 58. 最后一个单词的长度
// https://leetcode-cn.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	var res int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res == 0 {
				continue
			} else {
				break
			}
		}
		res++
	}
	return res
}
