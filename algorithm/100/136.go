// 136.只出现一次的数字
//给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
//
//
//
//
//
//
//
// 示例 1 ：
//
//
// 输入：nums = [2,2,1]
//
//
// 输出：1
//
// 示例 2 ：
//
//
// 输入：nums = [4,1,2,1,2]
//
//
// 输出：4
//
// 示例 3 ：
//
//
// 输入：nums = [1]
//
//
// 输出：1
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -3 * 10⁴ <= nums[i] <= 3 * 10⁴
// 除了某个元素只出现一次以外，其余每个元素均出现两次。
//
//
// Related Topics 位运算 数组 👍 3247 👎 0

package algorithm_100

// leetcode submit region begin(Prohibit modification and deletion)
// 1^0=1, 1^1=0, 任意数异或 0 等于本身，任意数与自身进行异或等于0
// 对数组各个数进行异或操作，相同的2个数异或后得到0，剩下的单个数异或后得到本身
func singleNumberV2(nums []int) int {
	var res int
	for _, n := range nums {
		res ^= n
	}
	return res
}

func singleNumberV3(nums []int) int {
	var res, tmp int16
	for i := 0; i < 16; i++ {
		tmp = 0
		for j := 0; j < len(nums); j++ {
			if int16(nums[j])&(1<<i) != 0 {
				tmp++
			}
		}

		if tmp&1 == 1 {
			res |= 1 << i
		}
	}
	return int(res)
}

//leetcode submit region end(Prohibit modification and deletion)
