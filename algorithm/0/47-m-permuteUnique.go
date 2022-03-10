package algorithm_0

import "sort"

// 47. 全排列 II
// https://leetcode-cn.com/problems/permutations-ii/

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var n, last = len(nums), -1
	var visited = make([]int, n)
	var arr []int

	var fn func(idx int)
	fn = func(idx int) {
		if idx == n {
			res = append(res, append([]int{}, arr...))
			return
		}
		last = -1
		for i := 0; i < n; i++ {
			if visited[i] == 1 || last >= 0 && nums[i] == nums[last] {
				continue
			}
			arr = append(arr, nums[i])
			visited[i] = 1
			fn(idx + 1)
			last = i
			arr = arr[:len(arr)-1]
			visited[i] = 0
		}
	}
	sort.Ints(nums)
	fn(0)
	return res
}

func permuteUnique1(nums []int) [][]int {
	var res [][]int
	var n, last = len(nums), -1
	var visited = make([]bool, n)

	var fn func(idx int, arr []int)
	fn = func(idx int, arr []int) {
		if len(arr)+1 == n {
			res = append(res, append(append([]int{}, arr...), nums[idx]))
			return
		}
		visited[idx] = true
		last = -1
		for i := 0; i < n; i++ {
			if visited[i] || (last >= 0 && nums[i] == nums[last]) {
				continue
			}
			fn(i, append(arr, nums[idx]))
			last = i
		}
		visited[idx] = false
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if last >= 0 && nums[i] == nums[last] {
			continue
		}
		fn(i, nil)
		last = i
	}
	return res
}
