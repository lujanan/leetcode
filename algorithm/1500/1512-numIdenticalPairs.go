package algorithm_1500

// 1512. 好数对的数目
// https://leetcode-cn.com/problems/number-of-good-pairs/

func numIdenticalPairs(nums []int) int {
	var nMap = make(map[int]int)
	for i := range nums {
		nMap[nums[i]]++
	}

	var res int
	for k := range nMap {
		if nMap[k] > 1 {
			res += nMap[k] * (nMap[k] - 1) / 2
		}
	}
	return res
}
