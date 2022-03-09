package algorithm_0

// 46. 全排列
// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	var res [][]int
	var n = len(nums)
	var visited = make([]bool, len(nums))
	var fn func(idx int, arr []int)
	fn = func(idx int, arr []int) {
		if len(arr) == n-1 {
			res = append(res, append(arr, nums[idx]))
			return
		}
		visited[idx] = true
		for i := 0; i < n; i++ {
			if !visited[i] {
				fn(i, append(arr, nums[idx]))
			}
		}
		visited[idx] = false
	}
	
	for i := 0; i < n; i++ {
		fn(i, nil)
	}
	return res
}
