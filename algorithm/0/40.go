// 40.组合总和II

package algorithm_0

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var ans = make([][]int, 0)
	sort.Ints(candidates)
	var dfs func(nums []int, total, idx int)
	dfs = func(nums []int, total, idx int) {
		if total >= target {
			if total == target {
				var tmp = make([]int, len(nums))
				copy(tmp, nums)
				ans = append(ans, tmp)
			}
			return
		}

		for i := idx; i < len(candidates); i++ {
			dfs(append(nums, candidates[i]), total+candidates[i], i+1)
			
			for ; i < len(candidates)-1; i++ {
				if candidates[i] != candidates[i+1] {
					break
				}
			}
		}
	}
	dfs([]int{}, 0, 0)

	return ans
}
