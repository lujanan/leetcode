package algorithm_300

// 387. 字符串中的第一个唯一字符
// https://leetcode-cn.com/problems/first-unique-character-in-a-string/

// 空间优化
func firstUniqChar(s string) int {
	var n = len(s)
	var pos = [26]int{}
	for i := 0; i < 26; i++ {
		pos[i] = n
	}

	for i := range s {
		if pos[s[i]-'a'] == n {
			pos[s[i]-'a'] = i
		} else {
			pos[s[i]-'a'] = n + 1
		}
	}
	var res = n
	for i := range pos {
		if pos[i] < res {
			res = pos[i]
		}
	}
	if res < n {
		return res
	}
	return -1
}

func firstUniqChar1(s string) int {
	var sMap = make(map[byte]int)
	for i := range s {
		sMap[s[i]]++
	}

	for i := range s {
		if sMap[s[i]] <= 1 {
			return i
		}
	}
	return -1
}
