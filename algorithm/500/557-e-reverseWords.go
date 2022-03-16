package algorithm_500

// 557. 反转字符串中的单词 III
// https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/

func reverseWords(s string) string {
	var arr = []byte(s)
	i, j := 0, 0
	for i < len(s) && j <= len(s) {
		j++
		if j == len(s) || arr[j] == ' ' {
			k := j - 1
			for i < k {
				if arr[i] == ' ' {
					i++
				} else {
					arr[i], arr[k] = arr[k], arr[i]
					i, k = i+1, k-1
				}
			}
			i = j + 1
		}
	}
	return string(arr)
}
