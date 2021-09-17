//一个 平方和三元组 (a,b,c) 指的是满足 a2 + b2 = c2 的 整数 三元组 a，b 和 c 。
//
// 给你一个整数 n ，请你返回满足 1 <= a, b, c <= n 的 平方和三元组 的数目。
//
//
//
// 示例 1：
//
// 输入：n = 5
//输出：2
//解释：平方和三元组为 (3,4,5) 和 (4,3,5) 。
//
//
// 示例 2：
//
// 输入：n = 10
//输出：4
//解释：平方和三元组为 (3,4,5)，(4,3,5)，(6,8,10) 和 (8,6,10) 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 250
//
// Related Topics 数学 枚举
// 👍 6 👎 0

package algorithm_1900

import (
	"fmt"
	"testing"
)

func Test_countTriples(t *testing.T) {
	var num = 765
	nv := float64(num) /100
	fmt.Println(nv)
	return

	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{5}, 2},
		{"t2", args{10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := countTriples(tt.args.n); gotRes != tt.wantRes {
				t.Errorf("countTriples() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
