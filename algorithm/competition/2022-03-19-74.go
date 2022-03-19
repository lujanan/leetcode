package algorithm_competition

import "container/heap"

func divideArray(nums []int) bool {
	var nMap = make(map[int]int)
	for i := range nums {
		nMap[nums[i]]++
	}
	var n = len(nums) >> 1
	if len(nMap) > n {
		return false
	}
	for _, v := range nMap {
		if v&1 != 0 {
			return false
		}
	}
	return true
}

func maximumSubsequenceCount(text string, pattern string) int64 {
	var p0, p1 = make([]int, 0), make([]int, 0)
	var s0, s1 int64
	for i := range text {
		if text[i] == pattern[0] && pattern[0] == pattern[1] {
			p0 = append(p0, 1)
			p1 = append(p1, 1)
			s0++
			s1++
		} else if text[i] == pattern[0] {
			s0++
			p0 = append(p0, 1)
			p1 = append(p1, 0)
		} else if text[i] == pattern[1] {
			s1++
			p1 = append(p1, 1)
			p0 = append(p0, 0)
		}
	}
	var n = len(p0)
	var ps1 = make([]int, n)
	for i := 0; i < n; i++ {
		ps1[i] += p1[i]
		if i > 0 {
			ps1[i] += ps1[i-1]
		}
	}

	var maxNum int64
	for i := 0; i < n; i++ {
		if p0[i] > 0 {
			maxNum += int64(ps1[n-1] - ps1[i])
		}
	}
	if s0 > s1 {
		maxNum += s0
	} else {
		maxNum += s1
	}
	return maxNum
}

func halveArray(nums []int) int {
	var res int
	var total float64
	var nh = new(NumHeap)
	heap.Init(nh)
	for i := range nums {
		total += float64(nums[i])
		heap.Push(nh, float64(nums[i]))
	}
	total = total / 2

	var del float64
	for del < total && nh.Len() > 0 {
		res++
		var num = heap.Pop(nh).(float64)
		num = num / 2
		del += num
		if num > 0 {
			heap.Push(nh, num)
		}
	}

	return res
}

type NumHeap []float64

func (h NumHeap) Len() int { return len(h) }
func (h NumHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h NumHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *NumHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *NumHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
