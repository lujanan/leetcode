//给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
// 请找到和最小的 k 个数对 (u1,v1), (u2,v2) ... (uk,vk) 。
//
//
//
// 示例 1:
//
//
//输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//输出: [1,2],[1,4],[1,6]
//解释: 返回序列中的前 3 对数：
//     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
//
// 示例 2:
//
//
//输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//输出: [1,1],[1,1]
//解释: 返回序列中的前 2 对数：
//     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//
//
// 示例 3:
//
//
//输入: nums1 = [1,2], nums2 = [3], k = 3
//输出: [1,3],[2,3]
//解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
//
//
//
//
// 提示:
//
//
// 1 <= nums1.length, nums2.length <= 10⁵
// -10⁹ <= nums1[i], nums2[i] <= 10⁹
// nums1 和 nums2 均为升序排列
// 1 <= k <= 1000
//
// Related Topics 数组 堆（优先队列） 👍 307 👎 0

package algorithm_300

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res [][]int
	var n, m = len(nums1), len(nums2)
	var h = &hp{nil, nums1, nums2}
	for i := 0; i < k && i < n; i++ {
		h.data = append(h.data, index{i, 0})
	}

	for h.Len() > 0 && len(res) < k {
		num := heap.Pop(h)
		res = append(res, []int{nums1[num.(index).i], nums2[num.(index).j]})
		if num.(index).j+1 < m {
			heap.Push(h, index{num.(index).i, num.(index).j + 1})
		}
	}
	return res
}

type index struct {
	i, j int
}
type hp struct {
	data         []index
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }

func (h hp) Less(i, j int) bool {
	return h.nums1[h.data[i].i]+h.nums2[h.data[i].j] < h.nums1[h.data[j].i]+h.nums2[h.data[j].j]
}
func (h hp) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *hp) Push(x interface{}) { h.data = append(h.data, x.(index)) }

func (h *hp) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}
