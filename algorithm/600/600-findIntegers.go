package algorithm_600

/**
给定一个正整数 n，找出小于或等于 n 的非负整数中，其二进制表示不包含 连续的1 的个数。
示例 1:

输入: 5
输出: 5
解释:
下面是带有相应二进制表示的非负整数<= 5：
0 : 0
1 : 1
2 : 10
3 : 11
4 : 100
5 : 101
其中，只有整数3违反规则（有两个连续的1），其他5个满足规则。

说明: 1 <= n <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


func findIntegers(n int) (count int) {
	dp := [31]int{1, 1}
    for i := 2; i < 31; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    for i, pre := 29, 0; i >= 0; i-- {
        val := 1 << i
        if n&val > 0 {
            count += dp[i+1]
            if pre == 1 {
                break
            }
            pre = 1
        } else {
            pre = 0
        }
        if i == 0 {
            count++
        }
    }
    return
}

// 暴力失败
func findIntegers1(n int) (count int) {
	for i := 0; i <= n; i++ {
		if (i>>1)&i > 0 || (i<<1)&i > 0 {
			continue
		}
		count++
	}
	return
}
