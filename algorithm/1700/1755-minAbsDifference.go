package algorithm_1700

import (
	"math"
	"sort"
)

// 1755. 最接近目标值的子序列和
// https://leetcode-cn.com/problems/closest-subsequence-sum/

func minAbsDifference(nums []int, goal int) int {
	var n = len(nums)
	var left = n >> 1
	var right = n - left

	var lNum, rNum = make([]int, 1<<left), make([]int, 1<<right)
	for i := 0; i < (1 << left); i++ {
		for j := 0; j < left; j++ {
			if (i>>j)&1 == 1 {
				lNum[i] += nums[j]
			}
		}
	}
	for i := 0; i < 1<<right; i++ {
		for j := 0; j < right; j++ {
			if (i>>j)&1 == 1 {
				rNum[i] += nums[left+j]
			}
		}
	}

	sort.Ints(lNum)
	sort.Ints(rNum)
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var res = math.MaxInt32
	for i := range lNum {
		res = min(res, int(math.Abs(float64(lNum[i]-goal))))
		if res == 0 {
			return res
		}
	}
	for i := range rNum {
		res = min(res, int(math.Abs(float64(rNum[i]-goal))))
		if res == 0 {
			return res
		}
	}

	for i, j := 0, len(rNum)-1; i < len(lNum) && j >= 0; {
		res = min(res, int(math.Abs(float64(lNum[i]+rNum[j]-goal))))
		if res == 0 {
			return res
		}
		if lNum[i]+rNum[j] > goal {
			j--
		} else {
			i++
		}
	}
	return res
}

func minAbsDifference1(nums []int, goal int) int {
	var n = len(nums)
	var half = n >> 1

	var lNum, rNum = []int{0}, []int{0}
	for i := 0; i < half; i++ {
		var ln = len(lNum)
		for j := 0; j < ln; j++ {
			lNum = append(lNum, lNum[j]+nums[i])
		}
	}
	for i := half; i < n; i++ {
		var rn = len(rNum)
		for j := 0; j < rn; j++ {
			rNum = append(rNum, rNum[j]+nums[i])
		}
	}

	sort.Ints(lNum)
	sort.Ints(rNum)
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var res = math.MaxInt32
	for i := range lNum {
		res = min(res, int(math.Abs(float64(lNum[i]-goal))))
		if res == 0 {
			return res
		}
	}
	for i := range rNum {
		res = min(res, int(math.Abs(float64(rNum[i]-goal))))
		if res == 0 {
			return res
		}
	}

	for i, j := 0, len(rNum)-1; i < len(lNum) && j >= 0; {
		res = min(res, int(math.Abs(float64(lNum[i]+rNum[j]-goal))))
		if res == 0 {
			return res
		}
		if lNum[i]+rNum[j] > goal {
			j--
		} else {
			i++
		}
	}
	return res
}
