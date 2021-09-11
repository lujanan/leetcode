//回旋镖定义为一组三个点，这些点各不相同且不在一条直线上。
//
// 给出平面上三个点组成的列表，判断这些点是否可以构成回旋镖。
//
//
//
// 示例 1：
//
// 输入：[[1,1],[2,3],[3,2]]
//输出：true
//
//
// 示例 2：
//
// 输入：[[1,1],[2,2],[3,3]]
//输出：false
//
//
//
// 提示：
//
//
// points.length == 3
// points[i].length == 2
// 0 <= points[i][j] <= 100
//
// Related Topics 几何 数学
// 👍 28 👎 0

package algorithm_1000

import "testing"

func Test_isBoomerang(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{[][]int{{1, 1}, {2, 3}, {3, 2}}}, true},
		{"t2", args{[][]int{{1, 1}, {2, 2}, {3, 3}}}, false},
		{"t3", args{[][]int{{1, 1}, {5, 5}, {7, 7}}}, false},
		{"t4", args{[][]int{{1, 1}, {1, 5}, {1, 7}}}, false},
		{"t5", args{[][]int{{1, 1}, {5, 1}, {7, 1}}}, false},
		{"t6", args{[][]int{{1, 1}, {5, 1}, {7, 2}}}, true},
		{"t7", args{[][]int{{0, 0}, {2, 1}, {2, 1}}}, false},
		{"t8", args{[][]int{{2, 1}, {1, 1}, {2, 1}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBoomerang(tt.args.points); got != tt.want {
				t.Errorf("isBoomerang() = %v, want %v", got, tt.want)
			}
		})
	}
}
