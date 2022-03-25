//给定一个整数 n ，返回 n! 结果中尾随零的数量。
//
// 提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：0
//解释：3! = 6 ，不含尾随 0
//
//
// 示例 2：
//
//
//输入：n = 5
//输出：1
//解释：5! = 120 ，有一个尾随 0
//
//
// 示例 3：
//
//
//输入：n = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁴
//
//
//
//
// 进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？
// Related Topics 数学 👍 628 👎 0

package algorithm_100

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"t1", args{0}, 0},
		// {"t2", args{3}, 0},
		// {"t3", args{5}, 1},
		{"t4", args{25}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
