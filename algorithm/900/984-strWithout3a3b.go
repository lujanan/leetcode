//给定两个整数 A 和 B，返回任意字符串 S，要求满足：
//
//
// S 的长度为 A + B，且正好包含 A 个 'a' 字母与 B 个 'b' 字母；
// 子串 'aaa' 没有出现在 S 中；
// 子串 'bbb' 没有出现在 S 中。
//
//
//
//
// 示例 1：
//
// 输入：A = 1, B = 2
//输出："abb"
//解释："abb", "bab" 和 "bba" 都是正确答案。
//
//
// 示例 2：
//
// 输入：A = 4, B = 1
//输出："aabaa"
//
//
//
// 提示：
//
//
// 0 <= A <= 100
// 0 <= B <= 100
// 对于给定的 A 和 B，保证存在满足要求的 S。
//
// Related Topics 贪心 字符串 👍 62 👎 0

package algorithm_900

func strWithout3a3b(a int, b int) (res string) {
	var ab = a >= b
	for a > 0 || b > 0 {
		if a > b {
			if a > 1 {
				res += "aa"
				a -= 2
			} else {
				res += "a"
				a--
			}
			if b > 0 {
				res += "b"
				b--
			}
		} else if b > a {
			if b > 1 {
				res += "bb"
				b -= 2
			} else {
				res += "b"
				b--
			}
			if a > 0 {
				res += "a"
				a--
			}
		} else if b == a {
			if ab {
				res += "ab"
			} else {
				res += "ba"
			}
			a--
			b--
		}
	}
	return
}
