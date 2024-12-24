//有一棵特殊的苹果树，一连 n 天，每天都可以长出若干个苹果。在第 i 天，树上会长出 apples[i] 个苹果，这些苹果将会在 days[i] 天后（也就
//是说，第 i + days[i] 天时）腐烂，变得无法食用。也可能有那么几天，树上不会长出新的苹果，此时用 apples[i] == 0 且 days[i] =
//= 0 表示。
//
// 你打算每天 最多 吃一个苹果来保证营养均衡。注意，你可以在这 n 天之后继续吃苹果。
//
// 给你两个长度为 n 的整数数组 days 和 apples ，返回你可以吃掉的苹果的最大数目。
//
//
//
// 示例 1：
//
// 输入：apples = [1,2,3,5,2], days = [3,2,1,4,2]
//输出：7
//解释：你可以吃掉 7 个苹果：
//- 第一天，你吃掉第一天长出来的苹果。
//- 第二天，你吃掉一个第二天长出来的苹果。
//- 第三天，你吃掉一个第二天长出来的苹果。过了这一天，第三天长出来的苹果就已经腐烂了。
//- 第四天到第七天，你吃的都是第四天长出来的苹果。
//
//
// 示例 2：
//
// 输入：apples = [3,0,0,0,0,2], days = [3,0,0,0,0,2]
//输出：5
//解释：你可以吃掉 5 个苹果：
//- 第一天到第三天，你吃的都是第一天长出来的苹果。
//- 第四天和第五天不吃苹果。
//- 第六天和第七天，你吃的都是第六天长出来的苹果。
//
//
//
//
// 提示：
//
//
// apples.length == n
// days.length == n
// 1 <= n <= 2 * 10⁴
// 0 <= apples[i], days[i] <= 2 * 10⁴
// 只有在 apples[i] = 0 时，days[i] = 0 才成立
//
//
// Related Topics 贪心 数组 堆（优先队列） 👍 232 👎 0

package algorithm_1700

import "container/heap"

// leetcode submit region begin(Prohibit modification and deletion)
func eatenApples(apples []int, days []int) int {
	var res, idx int
	var aHeap = new(AppleHeap)

	for idx = 0; idx < len(apples); idx++ {
		if apples[idx] > 0 {
			aHeap.Push(&AppleData{Num: apples[idx], Effect: idx, Expire: days[idx] + idx - 1})
		}
	}
	heap.Init(aHeap)

	idx = 0
	for aHeap.Len() > 0 {
		for aHeap.Len() > 0 {
			apple := heap.Pop(aHeap).(*AppleData)
			if apple != nil && apple.Num > 0 && apple.Expire >= idx {
				if apple.Effect < idx {
					heap.Push(aHeap, apple)
					break
				}
	
				res++
				if apple.Num > 1 && apple.Expire > idx {
					apple.Num -= 1
					heap.Push(aHeap, apple)
					break
				}
			}
		}
		idx++
	}
	return res
}

type AppleData struct {
	Num    int
	Expire int
	Effect int
}

type AppleHeap []*AppleData

func (h AppleHeap) Len() int {
	return len(h)
}

func (h AppleHeap) Less(i, j int) bool {
	return h[i].Expire < h[j].Expire
}

func (h AppleHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *AppleHeap) Push(x interface{}) {
	*h = append(*h, x.(*AppleData))
}

func (h *AppleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
