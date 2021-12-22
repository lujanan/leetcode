//对于字符串 S 和 T，只有在 S = T + ... + T（T 自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。
//
// 返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。
//
//
//
// 示例 1：
//
//
//输入：str1 = "ABCABC", str2 = "ABC"
//输出："ABC"
//
//
// 示例 2：
//
//
//输入：str1 = "ABABAB", str2 = "ABAB"
//输出："AB"
//
//
// 示例 3：
//
//
//输入：str1 = "LEET", str2 = "CODE"
//输出：""
//
//
//
//
// 提示：
//
//
// 1 <= str1.length <= 1000
// 1 <= str2.length <= 1000
// str1[i] 和 str2[i] 为大写英文字母
//
// Related Topics 数学 字符串 👍 218 👎 0

package algorithm_1000

import "testing"

func Test_gcdOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"ABCABC", "ABC"}, "ABC"},
		{"t2", args{"ABABAB", "ABAB"}, "AB"},
		{"t3", args{"LEET", "CODE"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
