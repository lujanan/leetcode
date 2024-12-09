//给你一个坐标 coordinates ，它是一个字符串，表示国际象棋棋盘中一个格子的坐标。下图是国际象棋棋盘示意图。
//
//
//
// 如果所给格子的颜色是白色，请你返回 true，如果是黑色，请返回 false 。
//
// 给定坐标一定代表国际象棋棋盘上一个存在的格子。坐标第一个字符是字母，第二个字符是数字。
//
//
//
// 示例 1：
//
//
//输入：coordinates = "a1"
//输出：false
//解释：如上图棋盘所示，"a1" 坐标的格子是黑色的，所以返回 false 。
//
//
// 示例 2：
//
//
//输入：coordinates = "h3"
//输出：true
//解释：如上图棋盘所示，"h3" 坐标的格子是白色的，所以返回 true 。
//
//
// 示例 3：
//
//
//输入：coordinates = "c7"
//输出：false
//
//
//
//
// 提示：
//
//
// coordinates.length == 2
// 'a' <= coordinates[0] <= 'h'
// '1' <= coordinates[1] <= '8'
//
//
// Related Topics 数学 字符串 👍 62 👎 0

package algorithm_1800

import "testing"

func Test_squareIsWhite(t *testing.T) {
	type args struct {
		coordinates string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{"a1"}, false},
		{"t2", args{"h3"}, true},
		{"t3", args{"c7"}, false},
		{"t4", args{"g6"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareIsWhite(tt.args.coordinates); got != tt.want {
				t.Errorf("squareIsWhite() = %v, want %v", got, tt.want)
			}
		})
	}
}
