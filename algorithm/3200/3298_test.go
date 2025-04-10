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

import "testing"

func Test_validSubstringCountV5(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"t1", args{"bcca", "abc"}, 1},
		{"t2", args{"abcabc", "abc"}, 10},
		{"t3", args{"abcabc", "aaabc"}, 0},
		{"t4", args{"bccagshda", "abcas"}, 1},
		{"t5", args{"ddccdddccdddccccdddccdcdcd", "ddc"}, 279},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validSubstringCountV5(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("validSubstringCountV5() = %v, want %v", got, tt.want)
			}
		})
	}
}
