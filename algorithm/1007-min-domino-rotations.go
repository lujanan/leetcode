//在一排多米诺骨牌中，A[i] 和 B[i] 分别代表第 i 个多米诺骨牌的上半部分和下半部分。（一个多米诺是两个从 1 到 6 的数字同列平铺形成的 ——
//该平铺的每一半上都有一个数字。）
//
// 我们可以旋转第 i 张多米诺，使得 A[i] 和 B[i] 的值交换。
//
// 返回能使 A 中所有值或者 B 中所有值都相同的最小旋转次数。
//
// 如果无法做到，返回 -1.
//
//
//
// 示例 1：
//
//
//
// 输入：A = [2,1,2,4,2,2], B = [5,2,6,2,3,2]
//输出：2
//解释：
//图一表示：在我们旋转之前， A 和 B 给出的多米诺牌。
//如果我们旋转第二个和第四个多米诺骨牌，我们可以使上面一行中的每个值都等于 2，如图二所示。
//
//
// 示例 2：
//
// 输入：A = [3,5,1,2,3], B = [3,6,3,3,4]
//输出：-1
//解释：
//在这种情况下，不可能旋转多米诺牌使一行的值相等。
//
//
//
//
// 提示：
//
//
// 1 <= A[i], B[i] <= 6
// 2 <= A.length == B.length <= 20000
//
// Related Topics 贪心 数组
// 👍 80 👎 0

package algorithm

func minDominoRotations(tops []int, bottoms []int) int {
	// var rt, rt0, rb, rb0 = 0, 0, 1, 1 // rt表示以 tops[0] 为底，rb 表示以 bottoms[0] 为底，rt0，rb0 则是交换旋转了第一个数，在以其为底
	var res = [4]int{0, 1, 0, 1}
	for i := 1; i < len(tops); i++ {
		if tops[i] != tops[0] && bottoms[i] != tops[0] {
			// tops[0] 为底无法满足旋转多米诺牌使一行的值相等
			res[0] = -1
			res[3] = -1
		}
		if bottoms[i] != bottoms[0] && tops[i] != bottoms[0] {
			// bottoms[0] 为底无法满足旋转多米诺牌使一行的值相等
			res[2] = -1
			res[1] = -1
		}
		if res[0] == -1 && res[2] == -1 {
			// top[0] 或 bottoms[0] 为底都无法满足相等值，则不可能旋转多米诺牌使一行的值相等
			return -1
		}

		if res[0] > -1 && tops[0] != tops[i] {
			res[0]++
		}
		if res[1] > -1 && bottoms[0] != tops[i] {
			res[1]++
		}
		if res[2] > -1 && bottoms[0] != bottoms[i] {
			res[2]++
		}
		if res[3] > -1 && tops[0] != bottoms[i] {
			res[3]++
		}
	}

	var n = -1
	for _, v := range res {
		if v >= 0 {
			if n < 0 {
				n = v
			} else if v < n {
				n = v
			}
		}
	}
	return n
}
