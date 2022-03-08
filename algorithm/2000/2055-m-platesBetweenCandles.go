//给你一个长桌子，桌子上盘子和蜡烛排成一列。给你一个下标从 0 开始的字符串 s ，它只包含字符 '*' 和 '|' ，其中 '*' 表示一个 盘子 ，'|
//' 表示一支 蜡烛 。
//
// 同时给你一个下标从 0 开始的二维整数数组 queries ，其中 queries[i] = [lefti, righti] 表示 子字符串 s[
//lefti...righti] （包含左右端点的字符）。对于每个查询，你需要找到 子字符串中 在 两支蜡烛之间 的盘子的 数目 。如果一个盘子在 子字符串中 左边和右边
// 都 至少有一支蜡烛，那么这个盘子满足在 两支蜡烛之间 。
//
//
// 比方说，s = "||**||**|*" ，查询 [3, 8] ，表示的是子字符串 "*||**|" 。子字符串中在两支蜡烛之间的盘子数目为 2 ，子字符
//串中右边两个盘子在它们左边和右边 都 至少有一支蜡烛。
//
//
// 请你返回一个整数数组 answer ，其中 answer[i] 是第 i 个查询的答案。
//
//
//
// 示例 1:
//
//
//
// 输入：s = "**|**|***|", queries = [[2,5],[5,9]]
//输出：[2,3]
//解释：
//- queries[0] 有两个盘子在蜡烛之间。
//- queries[1] 有三个盘子在蜡烛之间。
//
//
// 示例 2:
//
//
//
// 输入：s = "***|**|*****|**||**|*", queries = [[1,17],[4,5],[14,17],[5,11],[15,16
//]]
//输出：[9,0,0,0,0]
//解释：
//- queries[0] 有 9 个盘子在蜡烛之间。
//- 另一个查询没有盘子在蜡烛之间。
//
//
//
//
// 提示：
//
//
// 3 <= s.length <= 10⁵
// s 只包含字符 '*' 和 '|' 。
// 1 <= queries.length <= 10⁵
// queries[i].length == 2
// 0 <= lefti <= righti < s.length
//
// Related Topics 数组 字符串 二分查找 前缀和 👍 60 👎 0

package algorithm_2000

// 前缀和+预处理
func platesBetweenCandles(s string, queries [][]int) []int {
	var sum = make([]int, len(s))
	if s[0] == '*' {
		sum[0] = 1
	}
	for i := 1; i < len(s); i++ {
		sum[i] = sum[i-1]
		if s[i] == '*' {
			sum[i]++
		}
	}

	var left, right = make([]int, len(s)), make([]int, len(s))
	left[len(s)-1], right[0] = -1, -1
	for i := 1; i < len(s); i++ {
		if s[i] == '|' {
			right[i] = i
		} else {
			right[i] = right[i-1]
		}
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '|' {
			left[i] = i
		} else {
			left[i] = left[i+1]
		}
	}

	var res []int
	for _, q := range queries {
		if left[q[0]] >= 0 && right[q[1]] >= 0 && right[q[1]] >= left[q[0]] {
			res = append(res, sum[right[q[1]]]-sum[left[q[0]]])
		} else {
			res = append(res, 0)
		}
	}
	return res
}

// 前缀和+预处理
func platesBetweenCandles2(s string, queries [][]int) []int {
	var sum = make([][3]int, len(s))
	if s[0] == '*' {
		sum[0][0] = 1
	}
	for i := 1; i < len(s); i++ {
		sum[i] = sum[i-1]
		if s[i] == '*' {
			sum[i][0]++
		}
	}

	sum[len(s)-1][1], sum[0][2] = -1, -1
	for i := 1; i < len(s); i++ {
		if s[i] == '|' {
			sum[i][2] = i
		} else {
			sum[i][2] = sum[i-1][2]
		}
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '|' {
			sum[i][1] = i
		} else {
			sum[i][1] = sum[i+1][1]
		}
	}

	var res []int
	for _, q := range queries {
		if sum[q[0]][1] >= 0 && sum[q[1]][2] >= 0 && sum[q[1]][2] >= sum[q[0]][1] {
			res = append(res, sum[sum[q[1]][2]][0]-sum[sum[q[0]][1]][0])
		} else {
			res = append(res, 0)
		}
	}
	return res
}

// // 前缀和+二分查找
func platesBetweenCandles1(s string, queries [][]int) []int {
	var sum = make([]int, len(s))
	if s[0] == '*' {
		sum[0] = 1
	}
	for i := 1; i < len(s); i++ {
		sum[i] = sum[i-1]
		if s[i] == '*' {
			sum[i]++
		}
	}

	var mid, pn int
	var left = func(l, r int) int {
		for l <= r {
			if s[l] == '|' {
				return l
			}

			mid = l + (r-l)>>1
			pn = sum[mid] - sum[l]

			if s[mid] == '|' && pn+1 == mid-l {
				return mid
			}

			if (s[mid] == '|' && pn+1 < mid-l) || (s[mid] == '*' && pn < mid-l) {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return -1
	}
	var right = func(l, r int) int {
		for l <= r {
			if s[r] == '|' {
				return r
			}

			mid = l + (r-l)>>1
			pn = sum[r] - sum[mid]

			if s[mid] == '|' && pn == r-mid {
				return mid
			}

			if (s[mid] == '|' && pn < r-mid) || (s[mid] == '*' && pn < r-mid) {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return -1
	}

	var tl, tr int
	var res []int
	for _, q := range queries {
		tl, tr = left(q[0], q[1]), right(q[0], q[1])
		if tl >= 0 && tr >= 0 {
			res = append(res, sum[tr]-sum[tl])
		} else {
			res = append(res, 0)
		}
	}
	return res
}
