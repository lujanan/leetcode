package algorithm_200

import "strconv"

// 227. 基本计算器 II
// https://leetcode-cn.com/problems/basic-calculator-ii/

func calculate(s string) int {
	var num []int
	var tmpStr []byte
	var symbol byte

	var fn = func(symbol byte) {
		if len(tmpStr) <= 0 {
			return
		}

		tn, _ := strconv.Atoi(string(tmpStr))
		num = append(num, tn)
		tmpStr = tmpStr[:0]

		if (symbol == '*' || symbol == '/') && len(num) >= 2 {
			switch symbol {
			case '*':
				num[len(num)-2] = num[len(num)-2] * num[len(num)-1]
			case '/':
				num[len(num)-2] = num[len(num)-2] / num[len(num)-1]
			}
			num = num[:len(num)-1]
		} else if symbol == '-' {
			num[len(num)-1] = -num[len(num)-1]
		}
	}

	for i := range s {
		if s[i] == ' ' {
			continue
		}

		switch s[i] {
		case '+', '-', '*', '/':
			fn(symbol)
			symbol = s[i]
		default:
			tmpStr = append(tmpStr, s[i])
		}
	}
	fn(symbol)

	var res int
	for i := 0; i < len(num); i++ {
		res += num[i]
	}
	return res
}
