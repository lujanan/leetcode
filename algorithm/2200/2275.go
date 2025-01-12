// 2275. 按位与结果大于零的最长组合

package algorithm_2200

func largestCombination(candidates []int) int {
	var arr = make([]int, 25)
	var j int
	for i := 0; i < len(candidates); i++ {
		j = 0
		for candidates[i] > 0 {
			if candidates[i]&1 == 1 {
				arr[j]++
			}
			candidates[i] >>= 1
			j++
		}
	}

	j = 0
	for i := 0; i < len(arr); i++ {
		if j < arr[i] {
			j = arr[i]
		}
	}
	return j
}
