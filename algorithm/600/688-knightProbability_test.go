//在一个 n x n 的国际象棋棋盘上，一个骑士从单元格 (row, column) 开始，并尝试进行 k 次移动。行和列是 从 0 开始 的，所以左上单元格
//是 (0,0) ，右下单元格是 (n - 1, n - 1) 。
//
// 象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。
//
//
//
// 每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。
//
// 骑士继续移动，直到它走了 k 步或离开了棋盘。
//
// 返回 骑士在棋盘停止移动后仍留在棋盘上的概率 。
//
//
//
// 示例 1：
//
//
//输入: n = 3, k = 2, row = 0, column = 0
//输出: 0.0625
//解释: 有两步(到(1,2)，(2,1))可以让骑士留在棋盘上。
//在每一个位置上，也有两种移动可以让骑士留在棋盘上。
//骑士留在棋盘上的总概率是0.0625。
//
//
// 示例 2：
//
//
//输入: n = 1, k = 0, row = 0, column = 0
//输出: 1.00000
//
//
//
//
// 提示:
//
//
// 1 <= n <= 25
// 0 <= k <= 100
// 0 <= row, column <= n
//
// Related Topics 动态规划 👍 170 👎 0

package algorithm_600

import "testing"

func Test_knightProbability(t *testing.T) {
	type args struct {
		n      int
		k      int
		row    int
		column int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"t1", args{3, 2, 0, 0}, 0.0625},
		{"t2", args{1, 0, 0, 0}, 1.0000},
		{"t3", args{10, 8, 0, 0}, 0.04556},
		{"t4", args{8, 30, 6, 4}, 0.00019},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightProbability(tt.args.n, tt.args.k, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("knightProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
