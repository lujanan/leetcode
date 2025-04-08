// 1140.石子游戏II
//Alice 和 Bob 继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。
//
// Alice 和 Bob 轮流进行，Alice 先开始。最初，M = 1。
//
// 在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
//
// 游戏一直持续到所有石子都被拿走。
//
// 假设 Alice 和 Bob 都发挥出最佳水平，返回 Alice 可以得到的最大数量的石头。
//
//
//
// 示例 1：
//
//
//输入：piles = [2,7,9,4,4]
//输出：10
//解释：如果一开始 Alice 取了一堆，Bob 取了两堆，然后 Alice 再取两堆。Alice 可以得到 2 + 4 + 4 = 10 堆。
//如果 Alice 一开始拿走了两堆，那么 Bob 可以拿走剩下的三堆。在这种情况下，Alice 得到 2 + 7 = 9 堆。返回 10，因为它更大。
//
//
// 示例 2:
//
//
//输入：piles = [1,2,3,4,5,100]
//输出：104
//
//
//
//
// 提示：
//
//
// 1 <= piles.length <= 100
//
// 1 <= piles[i] <= 10⁴
//
//
// Related Topics 数组 数学 动态规划 博弈 前缀和 👍 315 👎 0

package algorithm_1100

import "testing"

func Test_stoneGameII(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{2, 7, 9, 4, 4}}, 10},
		{"t2", args{[]int{1, 2, 3, 4, 5, 100}}, 104},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameII(tt.args.piles); got != tt.want {
				t.Errorf("stoneGameII() = %v, want %v", got, tt.want)
			}
		})
	}
}
