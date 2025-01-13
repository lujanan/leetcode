// 201.数字范围按位与
//给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）
//。
//
//
//
// 示例 1：
//
//
//输入：left = 5, right = 7
//输出：4
//
//
// 示例 2：
//
//
//输入：left = 0, right = 0
//输出：0
//
//
// 示例 3：
//
//
//输入：left = 1, right = 2147483647
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= left <= right <= 2³¹ - 1
//
//
// Related Topics 位运算 👍 544 👎 0

package algorithm_200

// leetcode submit region begin(Prohibit modification and deletion)
// 找到 left, right 两个数字的二进制字符串的最长公共前缀即可
// 利用 n&(n-1) 消掉最右边的1，直至 left == right 即为最长公共前缀
func rangeBitwiseAnd(left int, right int) int {
	var res = right
	for res > left {
		res &= (res - 1)
	}
	return res
}

func rangeBitwiseAndV2(left int, right int) int {
	var res = left
	for i := left + 1; i <= right; i++ {
		res &= i
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
