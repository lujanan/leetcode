//给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
//
//
//
// 示例 1：
//
//
//输入：timePoints = ["23:59","00:00"]
//输出：1
//
//
// 示例 2：
//
//
//输入：timePoints = ["00:00","23:59","00:00"]
//输出：0
//
//
//
//
// 提示：
//
//
// 2 <= timePoints.length <= 2 * 10⁴
// timePoints[i] 格式为 "HH:MM"
//
// Related Topics 数组 数学 字符串 排序 👍 164 👎 0

package algorithm_500

import "testing"

func Test_findMinDifference(t *testing.T) {
	type args struct {
		timePoints []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{[]string{"23:59", "00:00"}}, 1},
		{"t2", args{[]string{"00:00", "23:59", "00:00"}}, 0},
		{"t3", args{[]string{"23:59", "10:00", "12:40", "15:00"}}, 140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findMinDifference(tt.args.timePoints); gotRes != tt.wantRes {
				t.Errorf("findMinDifference() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
