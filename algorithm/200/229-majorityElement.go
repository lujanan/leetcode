//给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
//
//
//
//
// 示例 1：
//
//
//输入：[3,2,3]
//输出：[3]
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：[1]
//
//
// 示例 3：
//
//
//输入：[1,1,1,3,3,2,2,2]
//输出：[1,2]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
// Related Topics 数组 哈希表 计数 排序 👍 465 👎 0

package algorithm_200

// 摩尔投票法：摩尔投票法的核心思想为对拼消耗。
// 首先我们考虑最基本的摩尔投票问题，比如找出一组数字序列中出现次数大于总数1/2的数字（并且假设这个数字一定存在）。
// 我们可以直接利用反证法证明这样的数字只可能有一个。
// 题目要求找出次数大于 n/3 的数，则最多有 2 个
func majorityElement(nums []int) (res []int) {
	var v1, v2, e1, e2 int
	for i := 0; i < len(nums); i++ {
		if v1 > 0 && e1 == nums[i] {
			v1++
		} else if v2 > 0 && e2 == nums[i] {
			v2++
		} else if v1 == 0 {
			e1 = nums[i]
			v1++
		} else if v2 == 0 {
			e2 = nums[i]
			v2++
		} else {
			v1--
			v2--
		}
	}

	var n1, n2 int
	for i := 0; i < len(nums); i++ {
		if v1 > 0 && nums[i] == e1 {
			n1++
		}
		if v2 > 0 && nums[i] == e2 {
			n2++
		}
	}
	if v1 > 0 && n1 > len(nums)/3 {
		res = append(res, e1)
	}
	if v2 > 0 && n2 > len(nums)/3 {
		res = append(res, e2)
	}
	return
}

func majorityElement0(nums []int) (res []int) {
	aMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := aMap[nums[i]]; !ok {
			aMap[nums[i]] = 0
		}
		aMap[nums[i]]++
	}
	for k, v := range aMap {
		if v > len(nums)/3 {
			res = append(res, k)
		}
	}
	return
}
