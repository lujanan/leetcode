// 1406.石子游戏III
//Alice 和 Bob 继续他们的石子游戏。几堆石子 排成一行 ，每堆石子都对应一个得分，由数组 stoneValue 给出。
//
// Alice 和 Bob 轮流取石子，Alice 总是先开始。在每个玩家的回合中，该玩家可以拿走剩下石子中的的前 1、2 或 3 堆石子 。比赛一直持续到所
//有石头都被拿走。
//
// 每个玩家的最终得分为他所拿到的每堆石子的对应得分之和。每个玩家的初始分数都是 0 。
//
// 比赛的目标是决出最高分，得分最高的选手将会赢得比赛，比赛也可能会出现平局。
//
// 假设 Alice 和 Bob 都采取 最优策略 。
//
// 如果 Alice 赢了就返回 "Alice" ，Bob 赢了就返回 "Bob"，分数相同返回 "Tie" 。
//
//
//
// 示例 1：
//
//
//输入：values = [1,2,3,7]
//输出："Bob"
//解释：Alice 总是会输，她的最佳选择是拿走前三堆，得分变成 6 。但是 Bob 的得分为 7，Bob 获胜。
//
//
// 示例 2：
//
//
//输入：values = [1,2,3,-9]
//输出："Alice"
//解释：Alice 要想获胜就必须在第一个回合拿走前三堆石子，给 Bob 留下负分。
//如果 Alice 只拿走第一堆，那么她的得分为 1，接下来 Bob 拿走第二、三堆，得分为 5 。之后 Alice 只能拿到分数 -9 的石子堆，输掉比赛。
//
//如果 Alice 拿走前两堆，那么她的得分为 3，接下来 Bob 拿走第三堆，得分为 3 。之后 Alice 只能拿到分数 -9 的石子堆，同样会输掉比赛。
//
//注意，他们都应该采取 最优策略 ，所以在这里 Alice 将选择能够使她获胜的方案。
//
// 示例 3：
//
//
//输入：values = [1,2,3,6]
//输出："Tie"
//解释：Alice 无法赢得比赛。如果她决定选择前三堆，她可以以平局结束比赛，否则她就会输。
//
//
//
//
// 提示：
//
//
// 1 <= stoneValue.length <= 5 * 10⁴
// -1000 <= stoneValue[i] <= 1000
//
//
// Related Topics 数组 数学 动态规划 博弈 👍 133 👎 0

package algorithm_1400

import "testing"

func Test_stoneGameIII(t *testing.T) {
	type args struct {
		stoneValue []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{[]int{1, 2, 3, 7}}, "Bob"},
		{"t2", args{[]int{1, 2, 3, -9}}, "Alice"},
		{"t3", args{[]int{1, 2, 3, 6}}, "Tie"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameIII(tt.args.stoneValue); got != tt.want {
				t.Errorf("stoneGameIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
