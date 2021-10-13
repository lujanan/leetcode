//给定一副牌，每张牌上都写着一个整数。
//
// 此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
//
//
// 每组都有 X 张牌。
// 组内所有的牌上都写着相同的整数。
//
//
// 仅当你可选的 X >= 2 时返回 true。
//
//
//
// 示例 1：
//
// 输入：[1,2,3,4,4,3,2,1]
//输出：true
//解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
//
//
// 示例 2：
//
// 输入：[1,1,1,2,2,2,3,3]
//输出：false
//解释：没有满足要求的分组。
//
//
// 示例 3：
//
// 输入：[1]
//输出：false
//解释：没有满足要求的分组。
//
//
// 示例 4：
//
// 输入：[1,1]
//输出：true
//解释：可行的分组是 [1,1]
//
//
// 示例 5：
//
// 输入：[1,1,2,2,2,2]
//输出：true
//解释：可行的分组是 [1,1]，[2,2]，[2,2]
//
//
//
//提示：
//
//
// 1 <= deck.length <= 10000
// 0 <= deck[i] < 10000
//
//
//
// Related Topics 数组 哈希表 数学 计数 数论 👍 236 👎 0

package algorithm_900

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	var deckMap = make(map[int]int)
	for _, v := range deck {
		if _, ok := deckMap[v]; ok {
			deckMap[v]++
		} else {
			deckMap[v] = 1
		}
	}
	var tmpNum = 0
	for _, v := range deckMap {
		if v < 2 {
			return false
		}
		if tmpNum <= 0 {
			tmpNum = v
			continue
		}
		tmpNum = gcd(v, tmpNum)
		if tmpNum < 2 {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
