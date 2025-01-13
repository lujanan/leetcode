// 201.数字范围按位与
//给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）
//。
//
//
//
// 示例 1：
//
//
//输入：left = 5, right = 7
//输出：4
//
//
// 示例 2：
//
//
//输入：left = 0, right = 0
//输出：0
//
//
// 示例 3：
//
//
//输入：left = 1, right = 2147483647
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= left <= right <= 2³¹ - 1
//
//
// Related Topics 位运算 👍 544 👎 0

package algorithm_200

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{5, 7}, 4},
		{"t2", args{4, 6}, 4},
		{"t3", args{1, 2147483647}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
