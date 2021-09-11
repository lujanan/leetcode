package algorithm_500

import "fmt"

func reverseStr(s string, k int) string {
	if len(s) < 2 {
		return s
	}
	var res []byte
	var idx int
	for {
		if idx+2*k > len(s) {
			break
		}
		for j := idx + k - 1; j >= idx; j-- {
			res = append(res, s[j])
		}
		res = append(res, s[idx+k:idx+2*k]...)
		idx = idx + 2*k
	}
	if idx < len(s) {
		if len(s)-idx < k {
			for j := len(s) - 1; j >= idx; j-- {
				res = append(res, s[j])
			}
		} else {
			for j := idx + k - 1; j >= idx; j-- {
				res = append(res, s[j])
			}
			res = append(res, s[idx+k:]...)
		}
	}
	fmt.Println(string(res))
	return string(res)
}
