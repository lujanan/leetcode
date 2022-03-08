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

func threeSum(nums []int) [][]int {

	var nMap = make(map[int]int)
	for _, v := range nums {
		nMap[v]++
	}
	var sortABC = func(a, b, c int) [3]int {
		if a > b {
			a, b = b, a
		}
		if a > c {
			a, c = c, a
		}
		if b > c {
			b, c = c, b
		}
		return [3]int{a, b, c}
	}
	var resMap = make(map[[3]int]int)

	var diff int
	var ok bool
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			diff = 0 - nums[i] - nums[j]
			if _, ok = nMap[diff]; ok {
				if diff == 0 && ((nums[i] == 0 && nMap[diff] > 2) || nums[i] != nums[j]) {
					resMap[sortABC(diff, nums[i], nums[j])] = 1
				} else if (nMap[diff] != nums[i] || (nMap[diff] == nums[i] && nMap[diff] > 1)) &&
					(nMap[diff] != nums[j] || (nMap[diff] == nums[j] && nMap[diff] > 1)) {
					resMap[sortABC(diff, nums[i], nums[j])] = 1
				}
			}
		}
	}

	var res [][]int
	for k := range resMap {
		res = append(res, k[:])
	}
	return res
}
