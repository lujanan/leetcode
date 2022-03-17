package algorithm_600

// 680. 验证回文字符串 Ⅱ
// https://leetcode-cn.com/problems/valid-palindrome-ii/

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			var lf, rf = true, true
			for i, j := l+1, r; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					lf = false
					break
				}
			}
			for i, j := l, r-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					rf = false
					break
				}
			}
			return lf || rf
		}
		l, r = l+1, r-1
	}
	return true
}
