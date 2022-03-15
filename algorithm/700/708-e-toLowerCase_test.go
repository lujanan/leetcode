//给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "Hello"
//输出："hello"
//
//
// 示例 2：
//
//
//输入：s = "here"
//输出："here"
//
//
// 示例 3：
//
//
//输入：s = "LOVELY"
//输出："lovely"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 100
// s 由 ASCII 字符集中的可打印字符组成
//
// Related Topics 字符串 👍 201 👎 0

package algorithm_700

import "testing"

func Test_toLowerCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"Hello"}, "hello"},
		{"t2", args{"here"}, "here"},
		{"t3", args{"LOVELY"}, "lovely"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.s); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
