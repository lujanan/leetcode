//三枚石子放置在数轴上，位置分别为 a，b，c。
//
// 每一回合，你可以从两端之一拿起一枚石子（位置最大或最小），并将其放入两端之间的任一空闲位置。形式上，假设这三枚石子当前分别位于位置 x, y, z 且
//x < y < z。那么就可以从位置 x 或者是位置 z 拿起一枚石子，并将该石子移动到某一整数位置 k 处，其中 x < k < z 且 k != y。
//
// 当你无法进行任何移动时，即，这些石子的位置连续时，游戏结束。
//
// 要使游戏结束，你可以执行的最小和最大移动次数分别是多少？ 以长度为 2 的数组形式返回答案：answer = [minimum_moves,
//maximum_moves]
//
//
//
// 示例 1：
//
//
//输入：a = 1, b = 2, c = 5
//输出：[1, 2]
//解释：将石子从 5 移动到 4 再移动到 3，或者我们可以直接将石子移动到 3。
//
//
// 示例 2：
//
//
//输入：a = 4, b = 3, c = 2
//输出：[0, 0]
//解释：我们无法进行任何移动。
//
//
//
//
// 提示：
//
//
// 1 <= a <= 100
// 1 <= b <= 100
// 1 <= c <= 100
// a != b, b != c, c != a
//
// Related Topics 脑筋急转弯 数学 👍 42 👎 0

package algorithm_1000

import (
	"reflect"
	"testing"
)

func Test_numMovesStones(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{1, 2, 5}, []int{1, 2}},
		{"t2", args{4, 3, 2}, []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := numMovesStones(tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("numMovesStones() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
