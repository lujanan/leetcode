// 2950.奇偶位数
//给你一个 正 整数 n 。
//
// 用 even 表示在 n 的二进制形式（下标从 0 开始）中值为 1 的偶数下标的个数。
//
// 用 odd 表示在 n 的二进制形式（下标从 0 开始）中值为 1 的奇数下标的个数。
//
// 请注意，在数字的二进制表示中，位下标的顺序 从右到左。
//
// 返回整数数组 answer ，其中 answer = [even, odd] 。
//
//
//
// 示例 1：
//
//
// 输入：n = 50
//
//
// 输出：[1,2]
//
// 解释：
//
// 50 的二进制表示是 110010。
//
// 在下标 1，4，5 对应的值为 1。
//
// 示例 2：
//
//
// 输入：n = 2
//
//
// 输出：[0,1]
//
// 解释：
//
// 2 的二进制表示是 10。
//
// 只有下标 1 对应的值为 1。
//
//
//
// 提示：
//
//
// 1 <= n <= 1000
//
//
// Related Topics 位运算 👍 29 👎 0

package algorithm_2900

// leetcode submit region begin(Prohibit modification and deletion)
func evenOddBit(n int) []int {
	var res = []int{0, 0}
	for i := 0; n > 0; {
		res[i] += n & 1
		n >>= 1
		i ^= 1
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
