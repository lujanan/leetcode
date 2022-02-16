//给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
//
// 进阶：不要 使用任何内置的库函数，如 sqrt 。
//
//
//
// 示例 1：
//
//
//输入：num = 16
//输出：true
//
//
// 示例 2：
//
//
//输入：num = 14
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= num <= 2^31 - 1
//
// Related Topics 数学 二分查找 👍 349 👎 0

package algorithm_300

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{4}, true},
		{"t2", args{8}, false},
		{"t3", args{9}, true},
		{"t4", args{17}, false},
		{"t5", args{100}, true},
		{"t6", args{130}, false},
		{"t7", args{82}, false},
		{"t8", args{625}, true},
		{"t9", args{1000}, false},
		{"t10", args{170}, false},
		{"t12", args{1}, true},
		{"t13", args{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
