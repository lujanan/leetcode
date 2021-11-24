//给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。
//
//
//
// 示例 1：
//
//
//输入：s = "owoztneoer"
//输出："012"
//
//
// 示例 2：
//
//
//输入：s = "fviefuro"
//输出："45"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s[i] 为 ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"] 这些字符之一
// s 保证是一个符合题目要求的字符串
//
// Related Topics 哈希表 数学 字符串 👍 145 👎 0

package algorithm_400

import "testing"

func Test_originalDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"owoztneoer"}, "012"},
		{"t2", args{"fviefuro"}, "45"},
		{"t3", args{"zezethreseeightvfothreeurenefoeighttwournineeightrotworo"}, "0022334478889"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := originalDigits(tt.args.s); got != tt.want {
				t.Errorf("originalDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
