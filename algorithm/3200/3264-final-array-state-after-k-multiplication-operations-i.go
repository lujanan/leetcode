//给你一个整数数组 nums ，一个整数 k 和一个整数 multiplier 。
//
// 你需要对 nums 执行 k 次操作，每次操作中：
//
//
// 找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。
// 将 x 替换为 x * multiplier 。
//
//
// 请你返回执行完 k 次乘运算之后，最终的 nums 数组。
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
//
//
// 示例 2：
//
//
// 输入：nums = [1,2], k = 3, multiplier = 4
//
//
// 输出：[16,8]
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
// [4, 2]
//
//
// 2 次操作后
// [4, 8]
//
//
// 3 次操作后
// [16, 8]
//
//
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100
// 1 <= k <= 10
// 1 <= multiplier <= 5
//
//
// Related Topics 数组 数学 模拟 堆（优先队列） 👍 19 👎 0

package algorithm_3200

import (
	"container/heap"
)

// leetcode submit region begin(Prohibit modification and deletion)
func getFinalState(nums []int, k int, multiplier int) []int {
	var state = new(HeapArray)
	for i, v := range nums {
		state.Push(&NumIdx{Num: v, Idx: i})
	}
	heap.Init(state)

	for i := 0; i < k; i++ {
		num := heap.Pop(state).(*NumIdx)
		num.Num *= multiplier
		heap.Push(state, num)
	}

	for i := 0; i < len(nums); i++ {
		num := state.Pop().(*NumIdx)
		nums[num.Idx] = num.Num
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)

type NumIdx struct {
	Num int
	Idx int
}

type HeapArray []*NumIdx

func (h HeapArray) Len() int {
	return len(h)
}

func (h HeapArray) Less(i, j int) bool {
	if h[i].Num == h[j].Num {
		return h[i].Idx < h[j].Idx
	}
	return h[i].Num < h[j].Num
}

func (h HeapArray) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapArray) Push(num interface{}) {
	*h = append(*h, num.(*NumIdx))
}

func (h *HeapArray) Pop() interface{} {
	var l = h.Len()
	old := *h
	num := old[l-1]
	*h = old[:l-1]
	return num
}
