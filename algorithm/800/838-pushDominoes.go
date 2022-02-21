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

func pushDominoes(dominoes string) string {
	// dp[i] = dp[i-1], dp[i+1]

	var str = []byte(dominoes)
	for i := 0; i < len(dominoes); i++ {
		if str[i] != '.' {
			continue
		}
		var left, right byte = '.', '.'
		if i-1 >= 0 {
			left = str[i-1]
		}

		var l, r = i, i + 1
		for r < len(dominoes) && str[r] == '.' {
			r++
		}
		if r < len(dominoes) {
			right = str[r]
		}
		i = r
		r--

		if left == 'R' && right == 'L' {
			for l < r {
				str[l], str[r] = 'R', 'L'
				l++
				r--
			}
		} else if left == 'R' {
			for l <= r {
				str[l] = 'R'
				l++
			}
		} else if right == 'L' {
			for l <= r {
				str[r] = 'L'
				r--
			}
		}
	}

	return string(str)
}
