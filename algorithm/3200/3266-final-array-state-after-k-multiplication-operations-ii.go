//给你一个整数数组 nums ，一个整数 k 和一个整数 multiplier 。 
//
// 你需要对 nums 执行 k 次操作，每次操作中： 
//
// 
// 找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。 
// 将 x 替换为 x * multiplier 。 
// 
//
// k 次操作以后，你需要将 nums 中每一个数值对 10⁹ + 7 取余。 
//
// 请你返回执行完 k 次乘运算以及取余运算之后，最终的 nums 数组。 
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [2,1,3,5,6], k = 5, multiplier = 2 
// 
//
// 输出：[8,4,6,5,6] 
//
// 解释： 
//
// 
// 
// 
// 操作 
// 结果 
// 
// 
// 1 次操作后 
// [2, 2, 3, 5, 6] 
// 
// 
// 2 次操作后 
// [4, 2, 3, 5, 6] 
// 
// 
// 3 次操作后 
// [4, 4, 3, 5, 6] 
// 
// 
// 4 次操作后 
// [4, 4, 6, 5, 6] 
// 
// 
// 5 次操作后 
// [8, 4, 6, 5, 6] 
// 
// 
// 取余操作后 
// [8, 4, 6, 5, 6] 
// 
// 
// 
//
// 示例 2： 
//
// 
// 输入：nums = [100000,2000], k = 2, multiplier = 1000000 
// 
//
// 输出：[999999307,999999993] 
//
// 解释： 
//
// 
// 
// 
// 操作 
// 结果 
// 
// 
// 1 次操作后 
// [100000, 2000000000] 
// 
// 
// 2 次操作后 
// [100000000000, 2000000000] 
// 
// 
// 取余操作后 
// [999999307, 999999993] 
// 
// 
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁴ 
// 1 <= nums[i] <= 10⁹ 
// 1 <= k <= 10⁹ 
// 1 <= multiplier <= 10⁶ 
// 
//
// Related Topics 数组 模拟 堆（优先队列） 👍 40 👎 0


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

	// 阶乘后比较
	// var ti, tj = h[i].Num, h[j].Num
	// if h[i].Times > 0 {
	// 	ti *= h[i].Times * tj
	// }
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
