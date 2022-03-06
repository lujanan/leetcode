package algorithm_400

// 442. 数组中重复的数据
// https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/

func findDuplicates(nums []int) []int {
	var tmp = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 {
			continue
		}
		tmp, nums[i] = nums[i]-1, 0
		for tmp >= 0 {
			if nums[tmp] > 0 {
				tmp, nums[tmp] = nums[tmp]-1, -1
			} else {
				tmp, nums[tmp] = -1, nums[tmp]-1
			}
		}
	}
	var res []int
	for i := range nums {
		if nums[i] == -2 {
			res = append(res, i+1)
		}
	}
	return res
}
