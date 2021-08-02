package algorithm

// 寻找数组的中心索引
// https://leetcode-cn.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	var left, right int
	for _, v := range nums {
		right += v
	}
	for i, v := range nums {
		if left == right-left-v {
			return i
		}
		left += v
	}
	return -1
}
