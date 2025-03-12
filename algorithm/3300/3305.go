// 3305.元音辅音字符串计数I
//给你一个字符串 word 和一个 非负 整数 k。
//
// 返回 word 的 子字符串 中，每个元音字母（'a'、'e'、'i'、'o'、'u'）至少 出现一次，并且 恰好 包含 k 个辅音字母的子字符串的总数。
//
//
//
//
// 示例 1：
//
//
// 输入：word = "aeioqq", k = 1
//
//
// 输出：0
//
// 解释：
//
// 不存在包含所有元音字母的子字符串。
//
// 示例 2：
//
//
// 输入：word = "aeiou", k = 0
//
//
// 输出：1
//
// 解释：
//
// 唯一一个包含所有元音字母且不含辅音字母的子字符串是 word[0..4]，即 "aeiou"。
//
// 示例 3：
//
//
// 输入：word = "ieaouqqieaouqq", k = 1
//
//
// 输出：3
//
// 解释：
//
// 包含所有元音字母并且恰好含有一个辅音字母的子字符串有：
//
//
// word[0..5]，即 "ieaouq"。
// word[6..11]，即 "qieaou"。
// word[7..12]，即 "ieaouq"。
//
//
//
//
// 提示：
//
//
// 5 <= word.length <= 250
// word 仅由小写英文字母组成。
// 0 <= k <= word.length - 5
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 19 👎 0

package algorithm_3300

// leetcode submit region begin(Prohibit modification and deletion)
func countOfSubstrings(word string, k int) int {
	var hit int
	var arr = make([]int, 6)

	var updateArr = func(sk int, c byte, isAdd bool) {
		var idx = 5
		switch c {
		case 'a':
			idx = 0
		case 'e':
			idx = 1
		case 'i':
			idx = 2
		case 'o':
			idx = 3
		case 'u':
			idx = 4
		}

		if isAdd {
			arr[idx]++
			if idx < 5 || arr[idx] >= sk {
				hit |= 1 << idx
			}

		} else {
			arr[idx]--
			if (idx < 5 && arr[idx] <= 0) || (idx == 5 && arr[idx] < sk) {
				hit &^= (1 << idx)
			}
		}
	}

	var countk = func(sk int) int {
		clear(arr)
		hit = 0

		var num int
		if sk == 0 {
			hit |= 1 << 5
		}

		for i, j := 0, 0; i <= j && j < len(word); j++ {
			updateArr(sk, word[j], true)
			for hit == 1<<6-1 {
				num += len(word) - j
				updateArr(sk, word[i], false)
				i++
			}
		}
		return num
	}

	// 滑动窗口求解
	// 直接获取辅音字母 =k 个的子串不好滑
	// 变相求解辅音字母 >=k 个的子串数量 num1=count(k) 和 >=k+1 个的子串数量 num2=count(k+1)
	// 那辅音字母 =k 个的子串数量 num=num1-num2
	return countk(k) - countk(k+1)
}

//leetcode submit region end(Prohibit modification and deletion)
