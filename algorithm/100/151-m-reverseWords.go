package algorithm_100

// 151. 颠倒字符串中的单词
// https://leetcode-cn.com/problems/reverse-words-in-a-string/

func reverseWordsV1(s string) string {
	var sbytes, n = []byte(s), len(s)
	var i = n - 1
	var j int
	var res = make([]byte, 0, n)

	for i >= 0 {
		for i >= 0 && sbytes[i] == ' ' {
			i--
			n--
		}

		j = i - 1
		for j >= 0 && sbytes[j] != ' ' {
			j--
		}

		if j+1 >= 0 && j+1 <= i {
			res = append(res, sbytes[j+1:i+1]...)
			res = append(res, ' ')
			n++
		}
		i = j
	}

	return string(res[:n-1])
}

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
