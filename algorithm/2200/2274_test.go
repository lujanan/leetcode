//Alice 管理着一家公司，并租用大楼的部分楼层作为办公空间。Alice 决定将一些楼层作为 特殊楼层 ，仅用于放松。
//
// 给你两个整数 bottom 和 top ，表示 Alice 租用了从 bottom 到 top（含 bottom 和 top 在内）的所有楼层。另给你一个
//整数数组 special ，其中 special[i] 表示 Alice 指定用于放松的特殊楼层。
//
// 返回不含特殊楼层的 最大 连续楼层数。
//
//
//
// 示例 1：
//
//
//输入：bottom = 2, top = 9, special = [4,6]
//输出：3
//解释：下面列出的是不含特殊楼层的连续楼层范围：
//- (2, 3) ，楼层数为 2 。
//- (5, 5) ，楼层数为 1 。
//- (7, 9) ，楼层数为 3 。
//因此，返回最大连续楼层数 3 。
//
//
// 示例 2：
//
//
//输入：bottom = 6, top = 8, special = [7,6,8]
//输出：0
//解释：每层楼都被规划为特殊楼层，所以返回 0 。
//
//
//
//
// 提示
//
//
// 1 <= special.length <= 10⁵
// 1 <= bottom <= special[i] <= top <= 10⁹
// special 中的所有值 互不相同
//
//
// Related Topics 数组 排序 👍 27 👎 0

package algorithm_2200

import "testing"

func Test_maxConsecutive(t *testing.T) {
	type args struct {
		bottom  int
		top     int
		special []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{2, 9, []int{4, 6}}, 3},
		{"t2", args{6, 8, []int{7, 6, 8}}, 0},
		{"t3", args{1, 8, []int{3, 7}}, 3},
		{"t4", args{1, 18, []int{3, 10, 9, 5, 7, 6, 2, 8, 4}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutive(tt.args.bottom, tt.args.top, tt.args.special); got != tt.want {
				t.Errorf("maxConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
