//n 张多米诺骨牌排成一行，将每张多米诺骨牌垂直竖立。在开始时，同时把一些多米诺骨牌向左或向右推。
//
// 每过一秒，倒向左边的多米诺骨牌会推动其左侧相邻的多米诺骨牌。同样地，倒向右边的多米诺骨牌也会推动竖立在其右侧的相邻多米诺骨牌。
//
// 如果一张垂直竖立的多米诺骨牌的两侧同时有多米诺骨牌倒下时，由于受力平衡， 该骨牌仍然保持不变。
//
// 就这个问题而言，我们会认为一张正在倒下的多米诺骨牌不会对其它正在倒下或已经倒下的多米诺骨牌施加额外的力。
//
// 给你一个字符串 dominoes 表示这一行多米诺骨牌的初始状态，其中：
//
//
// dominoes[i] = 'L'，表示第 i 张多米诺骨牌被推向左侧，
// dominoes[i] = 'R'，表示第 i 张多米诺骨牌被推向右侧，
// dominoes[i] = '.'，表示没有推动第 i 张多米诺骨牌。
//
//
// 返回表示最终状态的字符串。
//
//
// 示例 1：
//
//
//输入：dominoes = "RR.L"
//输出："RR.L"
//解释：第一张多米诺骨牌没有给第二张施加额外的力。
//
//
// 示例 2：
//
//
//输入：dominoes = ".L.R...LR..L.."
//输出："LL.RR.LLRRLL.."
//
//
//
//
// 提示：
//
//
// n == dominoes.length
// 1 <= n <= 10⁵
// dominoes[i] 为 'L'、'R' 或 '.'
//
// Related Topics 双指针 字符串 动态规划 👍 176 👎 0

package algorithm_800

import "bytes"

// BFS
func pushDominoes(dominoes string) string {
	var n = len(dominoes)
	var times = make([]int, n)    // 受力时间
	var force = make([][]byte, n) // 每个点的每次受力方向
	var q []int                   // 受力的点
	var res = make([]byte, n)     // 记录结果

	for i, v := range dominoes {
		times[i] = -1 // 初始化没有受力
		res[i] = '.'  // 初始化无受力状态
		if v != '.' {
			q = append(q, i)
			times[i] = 0                         // 初始受力
			force[i] = append(force[i], byte(v)) // 方向
		}
	}

	for len(q) > 0 {
		idx := q[0]
		q = q[1:]
		if len(force[idx]) > 1 {
			// 多个力作用，第一个力左右后不再改变 或者 多个力之间互相抵消
			continue
		}

		f := force[idx][0]
		res[idx] = f
		next := idx - 1
		if f == 'R' {
			next = idx + 1
		}

		// 边界值不处理
		if 0 > next || next >= n {
			continue
		}

		// 处理下一个牌
		curT := times[idx]
		if times[next] == -1 {
			q = append(q, next)
			times[next] = curT + 1
			force[next] = append(force[next], f)
		} else if times[next] == curT+1 {
			force[next] = append(force[next], f)
		}
	}

	return string(res)
}

func pushDominoes2(dominoes string) string {
	n := len(dominoes)
	q := []int{}
	time := make([]int, n)
	for i := range time {
		time[i] = -1
	}
	force := make([][]byte, n)
	for i, ch := range dominoes {
		if ch != '.' {
			q = append(q, i)
			time[i] = 0
			force[i] = append(force[i], byte(ch))
		}
	}

	ans := bytes.Repeat([]byte{'.'}, n)
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if len(force[i]) > 1 {
			continue
		}
		f := force[i][0]
		ans[i] = f
		ni := i - 1
		if f == 'R' {
			ni = i + 1
		}
		if 0 <= ni && ni < n {
			t := time[i]
			if time[ni] == -1 {
				q = append(q, ni)
				time[ni] = t + 1
				force[ni] = append(force[ni], f)
			} else if time[ni] == t+1 {
				force[ni] = append(force[ni], f)
			}
		}
	}
	return string(ans)
}

// 双指针 模拟
func pushDominoes1(dominoes string) string {

	var str = []byte(dominoes)
	var left = byte('L')
	for i := 0; i < len(dominoes); {
		if str[i] != '.' {
			left = str[i]
			i++
			continue
		}

		var l, r = i, i + 1
		for r < len(dominoes) && str[r] == '.' {
			r++
		}
		var right = byte('R')
		if r < len(dominoes) {
			right = str[r]
		}
		i = r

		if left == right {
			for l < r {
				str[l] = right
				l++
			}
		} else if left == 'R' && right == 'L' {
			r--
			for l < r {
				str[l], str[r] = left, right
				l++
				r--
			}
		}
		left = right
	}

	return string(str)
}
