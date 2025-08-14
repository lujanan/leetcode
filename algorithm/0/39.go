// 39.组合总和

package algorithm_0

func combinationSum(candidates []int, target int) [][]int {
	var ans = make([][]int, 0)
	var searchFn func(nums []int, total, idx int)
	searchFn = func(nums []int, total, idx int) {
		if total >= target {
			if total == target {
				var tmp = make([]int, len(nums))
				copy(tmp, nums)
				ans = append(ans, tmp)
			}
			return
		}

		for i := idx; i < len(candidates); i++ {
			searchFn(append(nums, candidates[i]), total+candidates[i], i)
		}
	}
	searchFn([]int{}, 0, 0)

	return ans
}
