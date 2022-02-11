package algorithm_1400

import "fmt"

// 暴力
func simplifiedFractions(n int) (res []string) {
	var fn func(a, b int) int
	fn = func(a, b int) int {
		if a%b == 0 {
			return b
		}
		return fn(b, a%b)
	}

	

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if fn(i, j) == 1 {
				res = append(res, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return
}
