package algorithm_200

// 242. 有效的字母异位词
// https://leetcode-cn.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var sMap, tMap = make(map[byte]int), make(map[byte]int)
	for i := range s {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	if len(sMap) != len(tMap) {
		return false
	}
	for k := range sMap {
		if _, ok := tMap[k]; !ok || sMap[k] != tMap[k] {
			return false
		}
	}
	return true
}
