//对于某些非负整数 k ，如果交换 s1 中两个字母的位置恰好 k 次，能够使结果字符串等于 s2 ，则认为字符串 s1 和 s2 的 相似度为 k 。
//
// 给你两个字母异位词 s1 和 s2 ，返回 s1 和 s2 的相似度 k 的最小值。
//
//
//
// 示例 1：
//
//
//输入：s1 = "ab", s2 = "ba"
//输出：1
//
//
// 示例 2：
//
//
//输入：s1 = "abc", s2 = "bca"
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= s1.length <= 20
// s2.length == s1.length
// s1 和 s2 只包含集合 {'a', 'b', 'c', 'd', 'e', 'f'} 中的小写字母
// s2 是 s1 的一个字母异位词
//
//
// Related Topics 广度优先搜索 字符串 👍 212 👎 0

package algorithm_800

import "math"

func kSimilarity(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	}

	var (
		bMap    = map[string]int{s2: 0}
		nextMap = make(map[string]int)
		minStep = math.MaxInt
	)

	for i := 0; i < len(s1); i++ {
		for str, num := range bMap {
			if len(str) == 1 {
				if minStep > num {
					minStep = num
				}
				continue
			}

			for k := 0; k < len(str); k++ {
				if str[k] != s1[i] {
					continue
				}

				var tb = append([]byte{}, str[1:]...)
				if k == 0 {
					if stepNum, ok := nextMap[string(tb)]; !ok || (ok && stepNum > num) {
						nextMap[string(tb)] = num
					}
					break
				}

				// 已经大于当前步骤次数，直接不要
				if num+1 > minStep {
					continue
				}

				tb[k-1] = str[0]
				if stepNum, ok := nextMap[string(tb)]; !ok || (ok && stepNum > num+1) {
					nextMap[string(tb)] = num + 1
				}
			}
		}
		bMap, nextMap = nextMap, make(map[string]int)

	}
	return minStep
}
