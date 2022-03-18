package algorithm_100

// 151. 颠倒字符串中的单词
// https://leetcode-cn.com/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
	var res []byte
	var i, j = len(s) - 1, len(s) - 1
	for i >= 0 && j >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i - 1
		for j >= 0 && s[j] != ' ' {
			j--
		}

		if j+1 >= 0 && j+1 <= i {
			res = append(res, s[j+1:i+1]...)
			res = append(res, ' ')
		}
		i = j
	}
	return string(res[:len(res)-1])
}
