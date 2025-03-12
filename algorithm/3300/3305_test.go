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

import "testing"

func Test_countOfSubstrings(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"aeioqq", 1}, 0},
		{"t2", args{"aeiou", 0}, 1},
		{"t3", args{"ieaouqqieaouqq", 1}, 3},
		{"t4", args{"iqeaouqi", 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfSubstrings(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("countOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
