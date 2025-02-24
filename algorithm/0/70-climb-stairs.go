//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 记忆化搜索 数学 动态规划
// 👍 1820 👎 0

package algorithm_0

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var dp = []int{1, 2}
	for i := 3; i <= n; i++ {
		dp[1], dp[0] = dp[1]+dp[0], dp[1]
	}
	return dp[1]
}

func climbStairsV3(n int) int {
	if n <= 2 {
		return n
	}

	var arr = []int{1, 2}
	for i := 3; i <= n; i++ {
		arr[0], arr[1] = arr[1], arr[0]+arr[1]
	}
	return arr[1]
}

func climbStairsV2(n int) int {
	if n <= 2 {
		return n
	}

	var p1, res = 1, 2
	for i := 3; i <= n; i++ {
		p1, res = res, res+p1
	}
	return res
}

func climbStairs1(n int) (res int) {
	if n <= 2 {
		return n
	}
	var pre1, pre2 = 1, 2
	for i := 3; i <= n; i++ {
		res = pre1 + pre2
		pre1 = pre2
		pre2 = res
	}
	return
}
