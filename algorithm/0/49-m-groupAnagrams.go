package algorithm_0

import "sort"

// 49. 字母异位词分组
// https://leetcode-cn.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	var strMap = make(map[string][]string)

	for i := range strs {
		var key = []byte(strs[i])
		sort.Slice(key, func(i, j int) bool {
			return key[i] <= key[j]
		})
		strMap[string(key)] = append(strMap[string(key)], strs[i])
	}

	var res [][]string
	for _, v := range strMap {
		res = append(res, v)
	}
	return res
}
