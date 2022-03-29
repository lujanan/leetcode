//一位老师正在出一场由 n 道判断题构成的考试，每道题的答案为 true （用 'T' 表示）或者 false （用 'F' 表示）。老师想增加学生对自己做出
//答案的不确定性，方法是 最大化 有 连续相同 结果的题数。（也就是连续出现 true 或者连续出现 false）。
//
// 给你一个字符串 answerKey ，其中 answerKey[i] 是第 i 个问题的正确结果。除此以外，还给你一个整数 k ，表示你能进行以下操作的最
//多次数：
//
//
// 每次操作中，将问题的正确答案改为 'T' 或者 'F' （也就是将 answerKey[i] 改为 'T' 或者 'F' ）。
//
//
// 请你返回在不超过 k 次操作的情况下，最大 连续 'T' 或者 'F' 的数目。
//
//
//
// 示例 1：
//
//
//输入：answerKey = "TTFF", k = 2
//输出：4
//解释：我们可以将两个 'F' 都变为 'T' ，得到 answerKey = "TTTT" 。
//总共有四个连续的 'T' 。
//
//
// 示例 2：
//
//
//输入：answerKey = "TFFT", k = 1
//输出：3
//解释：我们可以将最前面的 'T' 换成 'F' ，得到 answerKey = "FFFT" 。
//或者，我们可以将第二个 'T' 换成 'F' ，得到 answerKey = "TFFF" 。
//两种情况下，都有三个连续的 'F' 。
//
//
// 示例 3：
//
//
//输入：answerKey = "TTFTTFTT", k = 1
//输出：5
//解释：我们可以将第一个 'F' 换成 'T' ，得到 answerKey = "TTTTTFTT" 。
//或者我们可以将第二个 'F' 换成 'T' ，得到 answerKey = "TTFTTTTT" 。
//两种情况下，都有五个连续的 'T' 。
//
//
//
//
// 提示：
//
//
// n == answerKey.length
// 1 <= n <= 5 * 10⁴
// answerKey[i] 要么是 'T' ，要么是 'F'
// 1 <= k <= n
//
// Related Topics 字符串 二分查找 前缀和 滑动窗口 👍 106 👎 0

package algorithm_2000

import "sort"

// p[j] - p[i] <= k

// 前缀和 + 滑动窗口
func maxConsecutiveAnswers(answerKey string, k int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var res, tlSum, trSum, tl, flSum, frSum, fl int
	for r, v := range answerKey {
		if v == 'T' {
			trSum++
		} else {
			frSum++
		}

		for tlSum < trSum-k {
			if answerKey[tl] == 'T' {
				tlSum++
			}
			tl++
		}

		for flSum < frSum-k {
			if answerKey[fl] == 'F' {
				flSum++
			}
			fl++
		}

		res = max(res, max(r-tl, r-fl)+1)
	}

	return res
}

// 前缀和 + 二分查找
func maxConsecutiveAnswers1(answerKey string, k int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var pt, pf = make([]int, len(answerKey)+1), make([]int, len(answerKey)+1)
	for i, v := range answerKey {
		pt[i+1], pf[i+1] = pt[i], pf[i]
		if v == 'T' {
			pt[i+1]++
		} else {
			pf[i+1]++
		}
	}

	var res int
	for r := range pt {
		tl := sort.SearchInts(pt, pt[r]-k)
		fl := sort.SearchInts(pf, pf[r]-k)
		res = max(res, max(r-tl, r-fl))
	}
	return res
}
