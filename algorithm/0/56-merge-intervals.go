package algorithm_0

import "sort"

// 合并区间
// https://leetcode-cn.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}
	// 找到区间最大值
	var max = 0
	for _, v := range intervals {
		if v[1] > max {
			max = v[1]
		}
	}

	var (
		same    []int
		sameMap = make(map[int]bool)
		box     = make([]int, max+1)
	)
	for _, v := range intervals {
		if v[0] == v[1] { // 区间左边与右边相等
			sameMap[v[0]] = false
			continue
		}
		for i := v[0]; i <= v[1]; i++ {
			box[i]++
			if i == v[1] {
				box[i]--
			}
		}
	}
	for k := range sameMap {
		same = append(same, k)
	}
	if len(same) > 1 {
		sort.Ints(same)
	}

	var (
		result = make([][]int, 0)
		tmp    []int
		index  int
	)
	for i, v := range box {
		if v <= 0 && len(tmp) <= 0 {
			continue
		}

		if len(tmp) <= 0 {
			index = 0
			for _, v := range same {
				if v < i {
					result = append(result, []int{v, v})
					index++
				} else {
					break
				}
			}
			same = same[index:]
			tmp = append(tmp, i)

		} else if v == 0 && len(tmp) > 0 {
			tmp = append(tmp, i)
			index = 0
			for _, v := range same {
				if v >= tmp[0] && v <= tmp[1] {
					index++
				} else {
					break
				}
			}
			same = same[index:]
			result = append(result, []int{tmp[0], tmp[1]})
			tmp = tmp[:0]
		}
	}
	for _, v := range same {
		result = append(result, []int{v, v})
	}
	return result
}
