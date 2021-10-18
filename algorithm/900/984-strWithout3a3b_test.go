//给定两个整数 A 和 B，返回任意字符串 S，要求满足：
//
//
// S 的长度为 A + B，且正好包含 A 个 'a' 字母与 B 个 'b' 字母；
// 子串 'aaa' 没有出现在 S 中；
// 子串 'bbb' 没有出现在 S 中。
//
//
//
//
// 示例 1：
//
// 输入：A = 1, B = 2
//输出："abb"
//解释："abb", "bab" 和 "bba" 都是正确答案。
//
//
// 示例 2：
//
// 输入：A = 4, B = 1
//输出："aabaa"
//
//
//
// 提示：
//
//
// 0 <= A <= 100
// 0 <= B <= 100
// 对于给定的 A 和 B，保证存在满足要求的 S。
//
// Related Topics 贪心 字符串 👍 62 👎 0

package algorithm_900

import "testing"

func Test_strWithout3a3b(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{4, 1}, "aabaa"},
		{"t2", args{1, 2}, "bba"},
		{"t3", args{3, 3}, "ababab"},
		{"t4", args{1, 3}, "bbab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := strWithout3a3b(tt.args.a, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("strWithout3a3b() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
