package algorithm_0

// https://leetcode.cn/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	var res int
	for x > res {
		res = res*10 + x%10
		x /= 10
	}
	return res == x || x == res/10
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	var a, b int
	a = x
	for a > 0 {
		b = b*10 + a%10
		a /= 10
	}
	return b == x
}
