//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 4429 👎 0

package algorithm_0

import "sort"

func threeSumV2(nums []int) [][]int {
	var n = len(nums)
	sort.Ints(nums)
	var res [][]int
	for k := 0; k < n-2; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		for i, j := k+1, n-1; i < j; {
			var sum = nums[k] + nums[i] + nums[j]
			if sum == 0 {
				res = append(res, []int{nums[k], nums[i], nums[j]})
			}

			if sum <= 0 {
				i++
				for i < n && nums[i] == nums[i-1] {
					i++
				}

			} else {
				j--
				for j < n && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}

	return res
}

func threeSum(nums []int) [][]int {
	var n = len(nums)
	sort.Ints(nums)
	var res [][]int
	var k, i, j, sum int
	for k = 0; k < n-2; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		for i, j = k+1, n-1; i < n && j >= 0 && j > i; {
			sum = nums[k] + nums[i] + nums[j]
			if sum == 0 {
				res = append(res, []int{nums[k], nums[i], nums[j]})
			}

			if sum < 0 {
				i++
				for i < n && nums[i] == nums[i-1] {
					i++
				}
			}
			if sum >= 0 {
				j--
				for j >= 0 && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}

	return res
}
