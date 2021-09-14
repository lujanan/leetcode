//给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
//
// 如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
//输出："apple"
//
//
// 示例 2：
//
//
//输入：s = "abpcplea", dictionary = ["a","b","c"]
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// 1 <= dictionary.length <= 1000
// 1 <= dictionary[i].length <= 1000
// s 和 dictionary[i] 仅由小写英文字母组成
//
// Related Topics 数组 双指针 字符串 排序
// 👍 231 👎 0

package algorithm_500

import (
	"testing"
)

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// {"t1", args{"abpcplea", []string{"ale", "apple", "monkey", "plea"}}, "apple"},
		// {"t2", args{"abpcplea", []string{"a", "b", "c"}}, "a"},
		// {"t3", args{"abpcplea", []string{"aleq", "appleq", "moneyq", "pleaq"}}, ""},
		{"t4", args{"abce", []string{"abe", "abc"}}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findLongestWord(tt.args.s, tt.args.dictionary); gotRes != tt.wantRes {
				t.Errorf("findLongestWord() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
