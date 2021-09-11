//第 i 个人的体重为 people[i]，每艘船可以承载的最大重量为 limit。
//
// 每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
//
// 返回载到每一个人所需的最小船数。(保证每个人都能被船载)。
//
//
//
// 示例 1：
//
// 输入：people = [1,2], limit = 3
//输出：1
//解释：1 艘船载 (1, 2)
//
//
// 示例 2：
//
// 输入：people = [3,2,2,1], limit = 3
//输出：3
//解释：3 艘船分别载 (1, 2), (2) 和 (3)
//
//
// 示例 3：
//
// 输入：people = [3,5,3,4], limit = 5
//输出：4
//解释：4 艘船分别载 (3), (3), (4), (5)
//
// 提示：
//
//
// 1 <= people.length <= 50000
// 1 <= people[i] <= limit <= 30000
//
// Related Topics 贪心 数组 双指针 排序
// 👍 154 👎 0

package algorithm_800

import (
	"testing"
)

func Test_qSort(t *testing.T) {
	type args struct {
		people *[]int
		l      int
		r      int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{&[]int{1, 0, 2, 9, 3, 8, 4, 7, 5, 6}, 0, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qSort(tt.args.people, tt.args.l, tt.args.r)
		})
	}
}

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{"t1", args{[]int{3, 5, 3, 4}, 5}, 4},
		{"t2", args{[]int{3, 2, 2, 1}, 3}, 3},
		{"t3", args{[]int{1, 2}, 3}, 1},
		{"t4", args{[]int{3, 8, 7, 1, 4}, 9}, 3},
		{"t5", args{[]int{3, 8, 4, 9, 2, 2, 7, 1, 6, 10, 6, 7, 1, 7, 7, 6, 4, 4, 10, 1}, 10}, 11},
		{"t6", args{[]int{3, 2, 3, 2, 2}, 6}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNum := numRescueBoats(tt.args.people, tt.args.limit); gotNum != tt.wantNum {
				t.Errorf("numRescueBoats() = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}
