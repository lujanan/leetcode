package algorithm_0

// 循环
func subsets(nums []int) [][]int {
	var res = [][]int{{}}
	for _, n := range nums {
		l := len(res)
		for i := 0; i < l; i++ {
			res = append(res, append(append([]int{}, res[i]...), n))
		}
	}
	return res
}

// 递归
func subsets1(nums []int) [][]int {
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
