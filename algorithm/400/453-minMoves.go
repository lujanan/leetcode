//给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：3
//解释：
//只需要3次操作（注意每次操作会增加两个元素的值）：
//[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1]
//输出：0
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// 答案保证符合 32-bit 整数
//
// Related Topics 数组 数学 👍 309 👎 0

package algorithm_400

// tips: 每次操作既可以理解为使 n−1 个元素增加 1，也可以理解使 1 个元素减少 1。显然，后者更利于我们的计算
func minMoves(nums []int) (res int) {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		res += num - min
	}
	return
}
