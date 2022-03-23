//给定整数 n 和 k，返回 [1, n] 中字典序第 k 小的数字。
//
//
//
// 示例 1:
//
//
//输入: n = 13, k = 2
//输出: 10
//解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
//
//
// 示例 2:
//
//
//输入: n = 1, k = 1
//输出: 1
//
//
//
//
// 提示:
//
//
// 1 <= k <= n <= 10⁹
//
// Related Topics 字典树 👍 386 👎 0

package algorithm_400

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{1, 1}, 1},
		{"t2", args{13, 2}, 10},
		{"t3", args{10, 7}, 6},
		{"t4", args{567, 53}, 146},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
