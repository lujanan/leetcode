package algorithm_1300

import "sort"

func getKth(lo int, hi int, k int) int {
	var wMap = make(map[int]int)
	var wFn func(n int) int 
	wFn = func(n int) int {
		if n == 1 {
			return 0
		}
		if v, ok := wMap[n]; ok {
			return v
		}
		if n%2 == 0 {
			wMap[n] = wFn(n/2) + 1
			return wMap[n]
		}
		wMap[n] = wFn(3*n+1) + 1
		return wMap[n]
	}

	var arr = make([][2]int, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		arr = append(arr, [2]int{i,  wFn(i)})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		} 
		return arr[i][1] < arr[j][1]
	})
	return arr[k-1][0]
}


func getKthV2(lo int, hi int, k int) int {
	var wMap = make(map[int]int)
	var wFn func(n int) int
	wFn = func(n int) int {
		if v, ok := wMap[n]; ok {
			return v
		}
		if n == 1 {
			wMap[n] = 0
			return wMap[n]
		}
		if n&1 != 1 {
			wMap[n] = wFn(n/2) + 1
			return wMap[n]
		}
		wMap[n] = wFn(3*n+1) + 1
		return wMap[n]
	}

	var arr = make([]int, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		arr = append(arr, i)
	}
	sort.Slice(arr, func(i, j int) bool {
		if wFn(arr[i]) == wFn(arr[j]) {
			return arr[i] < arr[j]
		}
		return wFn(arr[i]) < wFn(arr[j])
	})
	return arr[k-1]
}
