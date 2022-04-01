//给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <= i < len(arr) / 2，都有 arr[2 * i
//+ 1] = 2 * arr[2 * i]” 时，返回 true；否则，返回 false。
//
//
//
// 示例 1：
//
//
//输入：arr = [3,1,3,6]
//输出：false
//
//
// 示例 2：
//
//
//输入：arr = [2,1,2,6]
//输出：false
//
//
// 示例 3：
//
//
//输入：arr = [4,-2,2,-4]
//输出：true
//解释：可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]
//
//
//
//
// 提示：
//
//
// 0 <= arr.length <= 3 * 10⁴
// arr.length 是偶数
// -10⁵ <= arr[i] <= 10⁵
//
// Related Topics 贪心 数组 哈希表 排序 👍 136 👎 0

package algorithm_900

import "sort"

func canReorderDoubled(arr []int) bool {
	var nums []int
	var nMap = make(map[int]int)
	for i := range arr {
		if _, ok := nMap[arr[i]]; !ok {
			nums = append(nums, arr[i])
		}
		nMap[arr[i]]++
	}
	// 单数 0
	if _, ok := nMap[0]; ok && nMap[0]&1 == 1 {
		return false
	}
	sort.Slice(nums, func(i, j int) bool {
		var a, b = nums[i], nums[j]
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		return a < b
	})

	for _, n := range nums {
		if nMap[n] <= 0 {
			continue
		}

		if _, ok := nMap[2*n]; !ok || nMap[2*n] < nMap[n] {
			return false
		}
		nMap[2*n] -= nMap[n]
	}
	return true
}

func canReorderDoubled1(arr []int) bool {
	sort.Ints(arr)
	var nMap = make(map[int]int)
	for i := range arr {
		nMap[arr[i]]++
	}

	for i := 0; i < len(arr); i++ {
		if nMap[arr[i]] <= 0 {
			continue
		}

		num := arr[i]
		if num < 0 {
			if num%2 != 0 {
				return false
			}
			if _, ok := nMap[num/2]; !ok || nMap[num/2] < 1 {
				return false
			}
			nMap[num/2]--
		} else {
			if _, ok := nMap[2*num]; !ok || nMap[2*num] < 1 {
				return false
			}
			nMap[2*num]--
		}
		nMap[num]--
	}
	return true
}
