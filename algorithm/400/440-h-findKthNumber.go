//给定整数 n 和 k，返回 [1, n] 中字典序第 k 小的数字。
//
//
//
// 示例 1:
//
//
//输入: n = 13, k = 2
//输出: 10
//解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
//
//
// 示例 2:
//
//
//输入: n = 1, k = 1
//输出: 1
//
//
//
//
// 提示:
//
//
// 1 <= k <= n <= 10⁹
//
// Related Topics 字典树 👍 386 👎 0

package algorithm_400

func getSteps(cur, n int) (steps int) {
	first, last := cur, cur
	for first <= n {
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return
}

func findKthNumber(n, k int) int {
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		if steps <= k {
			k -= steps
			cur++
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
