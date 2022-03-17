package algorithm_200

// 205. 同构字符串
// https://leetcode-cn.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	var sMap, tMap = make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := tMap[t[i]]; !ok {
			tMap[t[i]] = s[i]
		} else if tMap[t[i]] != s[i] {
			return false
		}

		if _, ok := sMap[s[i]]; !ok {
			sMap[s[i]] = t[i]
		} else if sMap[s[i]] != t[i] {
			return false
		}
	}
	return true
}
