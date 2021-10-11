//给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。
//
// 在比较时，字母是依序循环出现的。举个例子：
//
//
// 如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'
//
//
//
//
// 示例：
//
// 输入:
//letters = ["c", "f", "j"]
//target = "a"
//输出: "c"
//
//输入:
//letters = ["c", "f", "j"]
//target = "c"
//输出: "f"
//
//输入:
//letters = ["c", "f", "j"]
//target = "d"
//输出: "f"
//
//输入:
//letters = ["c", "f", "j"]
//target = "g"
//输出: "j"
//
//输入:
//letters = ["c", "f", "j"]
//target = "j"
//输出: "c"
//
//输入:
//letters = ["c", "f", "j"]
//target = "k"
//输出: "c"
//
//
//
//
// 提示：
//
//
// letters长度范围在[2, 10000]区间内。
// letters 仅由小写字母组成，最少包含两个不同的字母。
// 目标字母target 是一个小写字母。
//
// Related Topics 数组 二分查找 👍 128 👎 0

package algorithm_700

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// {"t1", args{[]byte{'c', 'f', 'j'}, 'a'}, 'c'},
		// {"t2", args{[]byte{'c', 'f', 'j'}, 'c'}, 'f'},
		// {"t3", args{[]byte{'c', 'f', 'j'}, 'd'}, 'f'},
		// {"t4", args{[]byte{'c', 'f', 'j'}, 'g'}, 'j'},
		// {"t5", args{[]byte{'c', 'f', 'j'}, 'j'}, 'c'},
		// {"t6", args{[]byte{'c', 'f', 'j'}, 'k'}, 'c'},
		{"t7", args{[]byte{'e', 'e', 'e', 'n', 'n', 'n'}, 'e'}, 'n'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
