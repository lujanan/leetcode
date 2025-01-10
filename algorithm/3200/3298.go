// 3298.统计重新排列后包含另一个字符串的子字符串数目II
//给你两个字符串 word1 和 word2 。
//
// 如果一个字符串 x 重新排列后，word2 是重排字符串的 前缀 ，那么我们称字符串 x 是 合法的 。
//
// 请你返回 word1 中 合法 子字符串 的数目。
//
// 注意 ，这个问题中的内存限制比其他题目要 小 ，所以你 必须 实现一个线性复杂度的解法。
//
//
//
// 示例 1：
//
//
// 输入：word1 = "bcca", word2 = "abc"
//
//
// 输出：1
//
// 解释：
//
// 唯一合法的子字符串是 "bcca" ，可以重新排列得到 "abcc" ，"abc" 是它的前缀。
//
// 示例 2：
//
//
// 输入：word1 = "abcabc", word2 = "abc"
//
//
// 输出：10
//
// 解释：
//
// 除了长度为 1 和 2 的所有子字符串都是合法的。
//
// 示例 3：
//
//
// 输入：word1 = "abcabc", word2 = "aaabc"
//
//
// 输出：0
//
//
//
// 解释：
//
//
// 1 <= word1.length <= 10⁶
// 1 <= word2.length <= 10⁴
// word1 和 word2 都只包含小写英文字母。
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 8 👎 0

package algorithm_3200

// leetcode submit region begin(Prohibit modification and deletion)
func validSubstringCountV5(word1 string, word2 string) int64 {
	var wordNum = make([]int, 26)
	for _, v := range word2 {
		wordNum[v-'a']--
	}

	var cnt int
	for _, v := range wordNum {
		if v < 0 {
			cnt++
		}
	}

	var update = func(idx, n int, cnt *int) {
		wordNum[idx] += n
		if n > 0 && wordNum[idx] == 0 {
			*cnt--
		} else if n < 0 && wordNum[idx] == -1 {
			*cnt++
		}
	}

	var res int64
	var i, j int
	for i < len(word1) {
		for j < len(word1) && cnt > 0 {
			update(int(word1[j]-'a'), 1, &cnt)
			j++
		}

		if cnt <= 0 {
			res += int64(len(word1) - j + 1)
		}

		update(int(word1[i]-'a'), -1, &cnt)
		i++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
