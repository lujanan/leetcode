package algorithm_1100

import "sort"

// 1122. 数组的相对排序
// https://leetcode-cn.com/problems/relative-sort-array/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var nMap = make([]int, 1001)
	for i := range arr1 {
		nMap[arr1[i]]++
	}

	var i int
	for _, v := range arr2 {
		for nMap[v] > 0 {
			arr1[i] = v
			i++
			nMap[v]--
		}
	}

	for k, v := range nMap {
		for v > 0 {
			arr1[i] = k
			i++
			v--
		}
	}
	return arr1
}

func relativeSortArray1(arr1 []int, arr2 []int) []int {
	var nMap = make(map[int]int)
	for i := range arr1 {
		nMap[arr1[i]]++
	}

	var i int
	for _, v := range arr2 {
		for nMap[v] > 0 {
			arr1[i] = v
			i++
			nMap[v]--
		}
		delete(nMap, v)
	}

	var left []int
	for k, v := range nMap {
		for v > 0 {
			left = append(left, k)
			v--
		}
	}
	sort.Ints(left)
	for j := range left {
		arr1[i] = left[j]
		i++
	}
	return arr1
}
