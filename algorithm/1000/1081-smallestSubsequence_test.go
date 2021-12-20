//返回 s 字典序最小的子序列，该子序列包含 s 的所有不同字符，且只包含一次。
//
// 注意：该题与 316 https://leetcode.com/problems/remove-duplicate-letters/ 相同
//
//
//
// 示例 1：
//
//
//输入：s = "bcabc"
//输出："abc"
//
//
// 示例 2：
//
//
//输入：s = "cbacdcbc"
//输出："acdb"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
// Related Topics 栈 贪心 字符串 单调栈 👍 131 👎 0

package algorithm_1000

import "testing"

func Test_smallestSubsequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"bcabc"}, "abc"},
		{"t2", args{"cbacdcbc"}, "acdb"},
		{"t3", args{"bcaskldhqowasdsadqwdkabc"}, "akldhosqwbc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubsequence(tt.args.s); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
