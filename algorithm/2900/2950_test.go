// 2950.奇偶位数
//给你一个 正 整数 n 。
//
// 用 even 表示在 n 的二进制形式（下标从 0 开始）中值为 1 的偶数下标的个数。
//
// 用 odd 表示在 n 的二进制形式（下标从 0 开始）中值为 1 的奇数下标的个数。
//
// 请注意，在数字的二进制表示中，位下标的顺序 从右到左。
//
// 返回整数数组 answer ，其中 answer = [even, odd] 。
//
//
//
// 示例 1：
//
//
// 输入：n = 50
//
//
// 输出：[1,2]
//
// 解释：
//
// 50 的二进制表示是 110010。
//
// 在下标 1，4，5 对应的值为 1。
//
// 示例 2：
//
//
// 输入：n = 2
//
//
// 输出：[0,1]
//
// 解释：
//
// 2 的二进制表示是 10。
//
// 只有下标 1 对应的值为 1。
//
//
//
// 提示：
//
//
// 1 <= n <= 1000
//
//
// Related Topics 位运算 👍 29 👎 0

package algorithm_2900

import (
	"reflect"
	"testing"
)

func Test_evenOddBit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{50}, []int{1, 2}},
		{"t2", args{2}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenOddBit(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evenOddBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
