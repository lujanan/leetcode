//冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
//
// 在加热器的加热半径范围内的每个房屋都可以获得供暖。
//
// 现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。
//
// 说明：所有供暖器都遵循你的半径标准，加热的半径也一样。
//
//
//
// 示例 1:
//
//
//输入: houses = [1,2,3], heaters = [2]
//输出: 1
//解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
//
//
// 示例 2:
//
//
//输入: houses = [1,2,3,4], heaters = [1,4]
//输出: 1
//解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
//
//
// 示例 3：
//
//
//输入：houses = [1,5], heaters = [2]
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= houses.length, heaters.length <= 3 * 10⁴
// 1 <= houses[i], heaters[i] <= 10⁹
//
// Related Topics 数组 双指针 二分查找 排序 👍 319 👎 0

package algorithm_400

import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	fn := func(list []int, n int, l, r int) int {
		for l <= r {
			m := l + (r-l)>>1
			if n == list[m] {
				return 0
			}
			if list[m] > n {
				if m-1 < 0 {
					return list[m] - n
				}
				if list[m-1] < n {
					if list[m]-n >= n-list[m-1] {
						return n - list[m-1]
					} else {
						return list[m] - n
					}
				}
				r = m - 1
			} else {
				if m+1 >= len(list) {
					return n - list[m]
				}
				if list[m+1] > n {
					if list[m+1]-n >= n-list[m] {
						return n - list[m]
					} else {
						return list[m+1] - n
					}
				}
				l = m + 1
			}
		}
		return -1
	}

	var min = 0
	for _, v := range houses {
		tmp := fn(heaters, v, 0, len(heaters)-1)
		if tmp > min {
			min = tmp
		}
	}
	return min
}