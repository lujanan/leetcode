package algorithm_1300

import "sort"

func minSetSize(arr []int) int {
	var numMap = make(map[int]int)
	for _, v := range arr {
		numMap[v]++
	}
	var res, cnt, half = 0, 0, len(arr) / 2
	var list = make([]int, 0, len(numMap))
	for _, v := range numMap {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool { return list[i] > list[j] })
	for i := 0; i < len(list); i++ {
		res++
		cnt += list[i]
		if cnt >= half {
			break
		}
	}
	return res
}
