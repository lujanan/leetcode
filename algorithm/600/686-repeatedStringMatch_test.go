//给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1。
//
// 注意：字符串 "abc" 重复叠加 0 次是 ""，重复叠加 1 次是 "abc"，重复叠加 2 次是 "abcabc"。
//
//
//
// 示例 1：
//
// 输入：a = "abcd", b = "cdabcdab"
//输出：3
//解释：a 重复叠加三遍后为 "abcdabcdabcd", 此时 b 是其子串。
//
//
// 示例 2：
//
// 输入：a = "a", b = "aa"
//输出：2
//
//
// 示例 3：
//
// 输入：a = "a", b = "a"
//输出：1
//
//
// 示例 4：
//
// 输入：a = "abc", b = "wxyz"
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= a.length <= 10⁴
// 1 <= b.length <= 10⁴
// a 和 b 由小写英文字母组成
//
// Related Topics 字符串 字符串匹配 👍 194 👎 0

package algorithm_600

import (
	"testing"
)

func Test_repeatedStringMatch(t *testing.T) {
	type args struct {
		a string
		b string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"abcd", "cdabcdab"}, 3},
		{"t2", args{"a", "aa"}, 2},
		{"t3", args{"a", "a"}, 1},
		{"t4", args{"abc", "wxyz"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("repeatedStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
