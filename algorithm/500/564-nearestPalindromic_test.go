//给定一个表示整数的字符串 n ，返回与它最近的回文整数（不包括自身）。如果不止一个，返回较小的那个。
//
// “最近的”定义为两个整数差的绝对值最小。
//
//
//
// 示例 1:
//
//
//输入: n = "123"
//输出: "121"
//
//
// 示例 2:
//
//
//输入: n = "1"
//输出: "0"
//解释: 0 和 2是最近的回文，但我们返回最小的，也就是 0。
//
//
//
//
// 提示:
//
//
// 1 <= n.length <= 18
// n 只由数字组成
// n 不含前导 0
// n 代表在 [1, 10¹⁸ - 1] 范围内的整数
//
// Related Topics 数学 字符串 👍 186 👎 0

package algorithm_500

import "testing"

func Test_nearestPalindromic(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {"t1", args{"123"}, "121"},
		// {"t2", args{"1"}, "0"},
		{"t3", args{"1002"}, "1001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestPalindromic(tt.args.n); got != tt.want {
				t.Errorf("nearestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
