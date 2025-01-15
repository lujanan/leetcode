// 224. 基本计算器

package algorithm_200

import (
	"strconv"
)

func calculateV2(s string) int {
	var nums []int
	var char []byte

	var cal = func() {
		var i, j = len(char) - 1, len(nums) - 1
		if i < 0 || j < 1 || char[i] == '(' {
			return
		}

		switch char[i] {
		case '+':
			nums[j-1] += nums[j]
		case '-':
			nums[j-1] -= nums[j]
		}
		char = char[:i]
		nums = nums[:j]
	}

	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}

		switch s[i] {
		case '(', '+', '-':
			if s[i] == '-' {
				var addZero bool
				j := i - 1
				for ; j >= 0; j-- {
					if s[j] == ' ' {
						continue
					}
					addZero = s[j] == '('
					break
				}
				if addZero || j == -1 {
					nums = append(nums, 0)
				}
			}
			char = append(char, s[i])
			i++

		case ')':
			char = char[:len(char)-1]
			cal()
			i++

		default:
			var numByte = make([]byte, 0)
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				numByte = append(numByte, s[i])
				i++
			}
			num, _ := strconv.ParseInt(string(numByte), 10, 64)
			nums = append(nums, int(num))
			cal()
		}
	}

	cal()
	return nums[0]
}
