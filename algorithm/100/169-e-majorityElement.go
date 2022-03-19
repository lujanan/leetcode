package algorithm_100

// 169. 多数元素
// https://leetcode-cn.com/problems/majority-element/

func majorityElement(nums []int) int {
	var nMap = make(map[int]int)
	for i := range nums {
		nMap[nums[i]]++
	}

	var n = len(nums)
	for i := range nMap {
		if nMap[i] > n/2 {
			return i
		}
	}
	return -1
}
