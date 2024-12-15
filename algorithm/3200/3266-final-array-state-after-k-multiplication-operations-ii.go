package algorithm_3200

import "container/heap"

func getFinalStateV2(nums []int, k int, multiplier int) []int {
	var state = new(FinalHeap)
	for i, v := range nums {
		state.Push(&NumHeap{v, 0, i})
	}
	heap.Init(state)

	for i := 0; i < k; i++ {
		num := heap.Pop(state).(*NumHeap)
		num.Times++
		heap.Push(state, num)
	}

	for i := 0; i < len(nums); i++ {
		num := heap.Pop(state).(*NumHeap)
		if num.Times > 0 {
			nums[num.Idx] = num.Num * num.Times * multiplier % 1000000007
		} else {
			nums[num.Idx] = num.Num % 1000000007
		}
	}
	return nums
}

type NumHeap struct {
	Num   int
	Times int
	Idx   int
}

type FinalHeap []*NumHeap

func (h FinalHeap) Len() int {
	return len(h)
}

func (h FinalHeap) Less(i, j int) bool {
	if h[i].Num == h[j].Num && h[i].Times == h[j].Times {
		return h[i].Idx < h[j].Idx
	}
	if h[i].Num < h[j].Num && h[i].Times < h[j].Times {
		return true
	}

	var ti, tj = h[i].Num, h[j].Num
	if h[i].Times > 0 {
		ti *= h[i].Times * 
	}
	return h[i].Num*h[i].Times < h[j].Num*h[j].Times
}

func (h FinalHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FinalHeap) Push(num interface{}) {
	*h = append(*h, num.(*NumHeap))
}

func (h *FinalHeap) Pop() interface{} {
	var l = h.Len()
	old := *h
	num := old[l-1]
	*h = old[:l-1]
	return num
}
