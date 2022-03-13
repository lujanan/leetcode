package algorithm_competition


// 6031. 找出数组中的所有 K 近邻下标
// https://leetcode-cn.com/problems/find-all-k-distant-indices-in-an-array/

func findKDistantIndices(nums []int, key int, k int) []int {
	var res []int
	var last = -1
	for j := range nums {
		if nums[j] == key {
			i := j - k
			if i <= last {
				i = last + 1
			}
			for ; i <= j+k; i++ {
				if i >= 0 && i < len(nums) && i > last {
					res = append(res, i)
					last = i
				}
			}
		}
	}
	return res
}
