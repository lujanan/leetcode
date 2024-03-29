//给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回 true ；否则返回 false 。
//
// 交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，接着交换 s[i] 和 s[j] 处的字符。
//
//
// 例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。
//
//
//
//
// 示例 1：
//
//
//输入：s = "ab", goal = "ba"
//输出：true
//解释：你可以交换 s[0] = 'a' 和 s[1] = 'b' 生成 "ba"，此时 s 和 goal 相等。
//
// 示例 2：
//
//
//输入：s = "ab", goal = "ab"
//输出：false
//解释：你只能交换 s[0] = 'a' 和 s[1] = 'b' 生成 "ba"，此时 s 和 goal 不相等。
//
// 示例 3：
//
//
//输入：s = "aa", goal = "aa"
//输出：true
//解释：你可以交换 s[0] = 'a' 和 s[1] = 'a' 生成 "aa"，此时 s 和 goal 相等。
//
//
// 示例 4：
//
//
//输入：s = "aaaaaaabc", goal = "aaaaaaacb"
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= s.length, goal.length <= 2 * 10⁴
// s 和 goal 由小写英文字母组成
//
// Related Topics 哈希表 字符串 👍 218 👎 0

package algorithm_800

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{"ab", "ba"}, true},
		{"t2", args{"ab", "ab"}, false},
		{"t3", args{"aa", "aa"}, true},
		{"t4", args{"abab", "baba"}, false},
		{"t5", args{"abcd", "bacd"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
