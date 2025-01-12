// 2275. 按位与结果大于零的最长组合

package algorithm_2200

// 数据范围 2^23 < 10^7 < 2^24 ,所以遍历24个位置上1的数量，
// 找到1数量最多的位置，便是最长的组合
func largestCombination(candidates []int) int {
	var res, tmp int
	for i := 0; i < 24; i++ {
		tmp = 0
		for j := 0; j < len(candidates); j++ {
			if candidates[j]&(1<<i) != 0 {
				tmp++
			}
		}
		if res < tmp {
			res = tmp
		}
	}
	return res
}

func largestCombinationV2(candidates []int) int {
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
