package algorithm_100

// 125. 验证回文串
// https://leetcode-cn.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	var l, r = 0, len(s) - 1
	var check = func(idx int) bool {
		if idx < 0 || idx >= len(s) {
			return false
		}
		return ('0' <= s[idx] && s[idx] <= '9') ||
			('a' <= s[idx] && s[idx] <= 'z') ||
			('A' <= s[idx] && s[idx] <= 'Z')
	}
	for l < r {
		for l < r && !check(l) {
			l++
		}
		for l < r && !check(r) {
			r--
		}

		if l != r {
			var lc, rc = s[l], s[r]
			if 'A' <= lc && lc <= 'Z' {
				lc = lc - 'A' + 'a'
			}
			if 'A' <= rc && rc <= 'Z' {
				rc = rc - 'A' + 'a'
			}
			if lc != rc {
				return false
			}
			l, r = l+1, r-1
		}
	}
	return true
}
