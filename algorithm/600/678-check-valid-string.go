package algorithm_600

/**
给定一个只包含三种字符的字符串：（ ，） 和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：

任何左括号 ( 必须有相应的右括号 )。
任何右括号 ) 必须有相应的左括号 ( 。
左括号 ( 必须在对应的右括号之前 )。
* 可以被视为单个右括号 ) ，或单个左括号 ( ，或一个空字符串。
一个空字符串也被视为有效字符串。

示例1：
输入: "()"
输出: True

示例2：
输入: "(*)"
输出: True

示例3：
输入: "(*))"
输出: True

注意:
字符串大小将在 [1，100] 范围内。
*/

// 栈解法
func checkValidString(s string) bool {
	var left, star []int
	for i, ch := range s {
		if ch == '(' {
			left = append(left, i)
		} else if ch == '*' {
			star = append(star, i)
		} else if len(left) > 0 {
			left = left[:len(left)-1]
		} else if len(star) > 0 {
			star = star[:len(star)-1]
		} else {
			return false
		}
	}

	if len(left) > 0 {
		if len(left) > len(star) {
			return false
		}
		for i := len(left) - 1; i >= 0; i-- {
			if left[i] > star[len(star)-1] {
				return false
			}
			star = star[:len(star)-1]
		}
	}
	return true
}
