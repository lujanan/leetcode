package algorithm_0

import "math"

// 8. 字符串转换整数 (atoi)
// https://leetcode-cn.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	var num, negative int64 = -1, 0
loop:
	for i := range s {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if num == -1 {
				num = int64(s[i] - '0')
				continue
			}
			num = num*10 + int64(s[i]-'0')
			if num >= math.MaxInt32 {
				break loop
			}
		case '-', '+':
			if negative == 0 && num == -1 {
				negative = 1
				if s[i] == '-' {
					negative = 2
				}
				continue
			}
			break loop
		default:
			if num >= 0 || negative > 0 {
				break loop
			} else if num == -1 && s[i] != ' ' {
				return 0
			}
		}
	}
	if num <= 0 {
		return 0
	}
	if negative == 2 {
		num = -num
	}

	if num <= math.MinInt32 {
		return math.MinInt32
	} else if num >= math.MaxInt32 {
		return math.MaxInt32
	}
	return int(num)
}
