//给你一个整数数组 nums ，你可以在 nums 上执行下述操作 任意次 ：
//
//
// 如果 gcd(nums[i], nums[j]) > 1 ，交换 nums[i] 和 nums[j] 的位置。其中 gcd(nums[i], nums[
//j]) 是 nums[i] 和 nums[j] 的最大公因数。
//
//
// 如果能使用上述交换方式将 nums 按 非递减顺序 排列，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
// 输入：nums = [7,21,3]
//输出：true
//解释：可以执行下述操作完成对 [7,21,3] 的排序：
//- 交换 7 和 21 因为 gcd(7,21) = 7 。nums = [21,7,3]
//- 交换 21 和 3 因为 gcd(21,3) = 3 。nums = [3,7,21]
//
//
// 示例 2：
//
// 输入：nums = [5,2,6,2]
//输出：false
//解释：无法完成排序，因为 5 不能与其他元素交换。
//
//
// 示例 3：
//
// 输入：nums = [10,5,9,3,15]
//输出：true
//解释：
//可以执行下述操作完成对 [10,5,9,3,15] 的排序：
//- 交换 10 和 15 因为 gcd(10,15) = 5 。nums = [15,5,9,3,10]
//- 交换 15 和 3 因为 gcd(15,3) = 3 。nums = [3,5,9,15,10]
//- 交换 10 和 15 因为 gcd(10,15) = 5 。nums = [3,5,9,10,15]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// 2 <= nums[i] <= 10⁵
//
// Related Topics 并查集 数组 数学 排序 👍 19 👎 0

package algorithm_1900

func gcdSort(nums []int) (res bool) {
	res = true
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				break
			}
			g := gcd(nums[j], nums[j-1])
			if g > 1 {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				return false
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
