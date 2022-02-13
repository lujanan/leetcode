package algorithm_100

// https://leetcode-cn.com/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var wMap = make(map[string]int)
	for _, v := range wordList {
		wMap[v] = 1
	}
	if _, ok := wMap[endWord]; !ok {
		return 0
	}
	var bMap, eMap = map[string]int{beginWord: 1}, map[string]int{endWord: 1}
	var next, check = make(map[string]int), make(map[string]int)
	var step = 1

loop:
	if len(bMap) > len(eMap) {
		bMap, eMap = eMap, bMap
	}
	for s := range bMap {
		for i := 0; i < len(s); i++ {
			for j := 0; j < 26; j++ {
				var temp = []byte(s)
				if temp[i] == byte(j)+'a' {
					continue
				}

				temp[i] = byte(j) + 'a'
				if _, ok := wMap[string(temp)]; ok {
					if _, ok := eMap[string(temp)]; ok {
						return step + 1
					}
					if _, ok := check[string(temp)]; ok {
						continue
					}

					next[string(temp)] = 1
					check[string(temp)] = 1
				}
			}
		}
	}
	if len(next) > 0 {
		step++
		bMap = next
		next = make(map[string]int)
		goto loop
	}

	return 0
}
