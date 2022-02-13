package algorithm_100

import "strings"

// https://leetcode-cn.com/problems/word-ladder-ii/

// 广度优先搜索， 各种剪枝
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var wMap = make(map[string]int)
	for _, v := range wordList {
		wMap[v] = 0
	}
	if _, ok := wMap[endWord]; !ok {
		return nil
	}

	var bMap = map[string]string{beginWord: beginWord}
	var res [][]string
	var check = make(map[string]int)
	var min = -1
	var step = 1

loop:
	var next = make(map[string]string)
	for link, s := range bMap {
		for i := 0; i < len(s); i++ {
			for j := 0; j < 26; j++ {
				var temp = []byte(s)
				if temp[i] == byte(j)+'a' {
					continue
				}

				temp[i] = byte(j) + 'a'
				if _, ok := wMap[string(temp)]; !ok {
					continue
				}

				if string(temp) == endWord && (min == -1 || step+1 <= min) {
					if step+1 < min {
						res = res[:0]
					}
					var r1 = link + "," + string(temp)
					res = append(res, strings.Split(r1, ","))
					min = step + 1
				}

				if _, ok := check[string(temp)]; ok && check[string(temp)] < step+1 {
					continue
				}

				check[string(temp)] = step + 1
				next[link+","+string(temp)] = string(temp)
			}
		}
	}

	if len(next) > 0 {
		step++
		bMap = next
		goto loop
	}

	return res
}
