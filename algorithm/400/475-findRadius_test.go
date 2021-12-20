//冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
//
// 在加热器的加热半径范围内的每个房屋都可以获得供暖。
//
// 现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。
//
// 说明：所有供暖器都遵循你的半径标准，加热的半径也一样。
//
//
//
// 示例 1:
//
//
//输入: houses = [1,2,3], heaters = [2]
//输出: 1
//解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
//
//
// 示例 2:
//
//
//输入: houses = [1,2,3,4], heaters = [1,4]
//输出: 1
//解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
//
//
// 示例 3：
//
//
//输入：houses = [1,5], heaters = [2]
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= houses.length, heaters.length <= 3 * 10⁴
// 1 <= houses[i], heaters[i] <= 10⁹
//
// Related Topics 数组 双指针 二分查找 排序 👍 319 👎 0

package algorithm_400

import "testing"

func Test_findRadius(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"t1", args{[]int{1, 2, 3}, []int{2}}, 1},
		// {"t2", args{[]int{1, 2, 3, 4}, []int{1, 4}}, 1},
		// {"t3", args{[]int{1, 5}, []int{2}}, 3},
		// {"t4", args{[]int{1, 2, 3, 17, 25, 38, 50}, []int{2, 5, 30, 70}}, 20},
		// {"t5", args{[]int{10, 23, 45, 23, 45, 67, 89, 123, 34, 268}, []int{89, 34, 12, 45, 67, 123, 657, 509}}, 145},
		// {"t6", args{[]int{1}, []int{1, 2, 3, 4}}, 0},
		{"t7", args{[]int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923}, []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}}, 161834419},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.houses, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
