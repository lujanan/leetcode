//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
//
// Related Topics 栈 字符串 动态规划 👍 1685 👎 0

package algorithm_0

// 栈
func longestValidParentheses(s string) int {
	var stack = []int{-1}
	var res int
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) < 1 {
			stack = append(stack, i)
			continue
		}

		if res < i-stack[len(stack)-1] {
			res = i - stack[len(stack)-1]
		}
	}
	return res
}

// 动态规划2
func longestValidParentheses2(s string) int {
	var res int
	var dp = make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] != ')' {
			continue
		}

		if s[i-1] == '(' || (i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(') {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}

		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

// 动态规划1
func longestValidParentheses1(s string) int {
	var res, left int
	var dp = make([]int, len(s)+1)
	for i := range s {
		if s[i] == '(' {
			left++
			continue
		}

		if left > 0 {
			dp[i+1] = dp[i] + 2
			dp[i+1] += dp[i+1-dp[i+1]]
			left--
			if res < dp[i+1] {
				res = dp[i+1]
			}
		}
	}

	return res
}

// 双向计数
// 从左往右计数, left 记录'(' 数量, right 记录 ')' 数量，right == left 时，记录当前长度，right > left 时，清0
// 从右往左计数, left 记录'(' 数量, right 记录 ')' 数量，right == left 时，记录当前长度，right < left 时，清0
// 单向计数无法处理 ((()、())) 这些情况，所以需要双向计数
func longestValidParentheses3(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if maxLength < right*2 {
				maxLength = right * 2
			}
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if maxLength < left*2 {
				maxLength = left * 2
			}
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}
