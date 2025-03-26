// 2829.k-avoiding数组的最小总和
//给你两个整数 n 和 k 。
//
// 对于一个由 不同 正整数组成的数组，如果其中不存在任何求和等于 k 的不同元素对，则称其为 k-avoiding 数组。
//
// 返回长度为 n 的 k-avoiding 数组的可能的最小总和。
//
//
//
// 示例 1：
//
//
//输入：n = 5, k = 4
//输出：18
//解释：设若 k-avoiding 数组为 [1,2,4,5,6] ，其元素总和为 18 。
//可以证明不存在总和小于 18 的 k-avoiding 数组。
//
//
// 示例 2：
//
//
//输入：n = 2, k = 6
//输出：3
//解释：可以构造数组 [1,2] ，其元素总和为 3 。
//可以证明不存在总和小于 3 的 k-avoiding 数组。
//
//
//
//
// 提示：
//
//
// 1 <= n, k <= 50
//
//
// Related Topics 贪心 数学 👍 40 👎 0

package algorithm_2800

// leetcode submit region begin(Prohibit modification and deletion)

func minimumSum(n int, k int) int {
	var res int

	for i, j := 1, 0; j < n; i++ {
		if i > k/2 && i < k {
			continue
		}

		res += i
		j++
	}
	return res
}

func minimumSumV2(n int, k int) int {
	var res int
	var outNum = make(map[int]bool)

	for i, j := 1, 0; j < n; i++ {
		if _, ok := outNum[i]; ok {
			continue
		}
		res += i
		if k > i {
			outNum[k-i] = true
		}

		j++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
