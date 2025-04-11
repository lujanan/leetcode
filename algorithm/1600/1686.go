// 1686.石子游戏6
//Alice 和 Bob 轮流玩一个游戏，Alice 先手。
//
// 一堆石子里总共有 n 个石子，轮到某个玩家时，他可以 移出 一个石子并得到这个石子的价值。Alice 和 Bob 对石子价值有 不一样的的评判标准 。双方
//都知道对方的评判标准。
//
// 给你两个长度为 n 的整数数组 aliceValues 和 bobValues 。aliceValues[i] 和 bobValues[i] 分别表示
//Alice 和 Bob 认为第 i 个石子的价值。
//
// 所有石子都被取完后，得分较高的人为胜者。如果两个玩家得分相同，那么为平局。两位玩家都会采用 最优策略 进行游戏。
//
// 请你推断游戏的结果，用如下的方式表示：
//
//
// 如果 Alice 赢，返回 1 。
// 如果 Bob 赢，返回 -1 。
// 如果游戏平局，返回 0 。
//
//
//
//
// 示例 1：
//
//
//输入：aliceValues = [1,3], bobValues = [2,1]
//输出：1
//解释：
//如果 Alice 拿石子 1 （下标从 0开始），那么 Alice 可以得到 3 分。
//Bob 只能选择石子 0 ，得到 2 分。
//Alice 获胜。
//
//
// 示例 2：
//
//
//输入：aliceValues = [1,2], bobValues = [3,1]
//输出：0
//解释：
//Alice 拿石子 0 ， Bob 拿石子 1 ，他们得分都为 1 分。
//打平。
//
//
// 示例 3：
//
//
//输入：aliceValues = [2,4,3], bobValues = [1,6,7]
//输出：-1
//解释：
//不管 Alice 怎么操作，Bob 都可以得到比 Alice 更高的得分。
//比方说，Alice 拿石子 1 ，Bob 拿石子 2 ， Alice 拿石子 0 ，Alice 会得到 6 分而 Bob 得分为 7 分。
//Bob 会获胜。
//
//
//
//
// 提示：
//
//
// n == aliceValues.length == bobValues.length
// 1 <= n <= 10⁵
// 1 <= aliceValues[i], bobValues[i] <= 100
//
//
// Related Topics 贪心 数组 数学 博弈 排序 堆（优先队列） 👍 125 👎 0

package algorithm_1600

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func stoneGameVI(aliceValues []int, bobValues []int) int {
	// 取最大分值，取自己的或对方的，
	// 把对方的最大分值取走，虽然自己获得的分值减小，但对方获得的分值也减小了
	var abArr = make([][2]int, len(aliceValues))
	for i := 0; i < len(aliceValues); i++ {
		abArr[i] = [2]int{aliceValues[i] + bobValues[i], i}
	}
	sort.Slice(abArr, func(i, j int) bool { return abArr[i][0] > abArr[j][0] })

	var res int
	var isAlice = true
	for i := 0; i < len(aliceValues); i++ {
		if isAlice {
			res += aliceValues[abArr[i][1]]
		} else {
			res -= bobValues[abArr[i][1]]
		}
		isAlice = !isAlice
	}

	if res > 0 {
		return 1
	} else if res < 0 {
		return -1
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
