package algorithm_0

func subsets(nums []int) [][]int {
	var res [][]int
	var fn func(sub []int, idx int)
	fn = func(sub []int, idx int) {
		if idx >= len(nums) {
			res = append(res, sub)
			return
		}

		fn(append([]int{}, sub...), idx+1)
		fn(append(sub, nums[idx]), idx+1)
	}

	fn(nil, 0)
	return res
}
