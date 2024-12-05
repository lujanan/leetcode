package algorithm_0

func generateParenthesis(n int) (res []string) {
	var gen = func(left, right int, str string) {}
	gen = func(left, right int, str string) {
		if left == n && right == n {
			res = append(res, str)
			return
		}
		if left < n {
			gen(left+1, right, str+"(")
		}
		if left > right {
			gen(left, right+1, str+")")
		}
	}

	gen(0, 0, "")
	return res
}

func generateParenthesisV2(n int) (res []string) {

	var gen func(left, right int, str string)
	gen = func(left, right int, str string) {
		if n == left && n == right {
			res = append(res, str)
			return
		}

		if left < n {
			gen(left+1, right, str+"(")
		}
		if right < n && right < left {
			gen(left, right+1, str+")")
		}
	}

	gen(0, 0, "")
	return res
}
