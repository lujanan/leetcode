package algorithm_0

// 加一
// https://leetcode-cn.com/problems/plus-one/
func plusOne1(digits []int) []int {
	var length = len(digits)
	for i := 0; i < length/2; i++ {
		digits[i], digits[length-i-1] = digits[length-i-1], digits[i]
	}
	var add = 1
	for i, v := range digits {
		digits[i] = v + add
		if digits[i] > 9 {
			digits[i] = 0
		} else {
			add = 0
		}
	}
	if add > 0 {
		digits = append(digits, 1)
	}
	length = len(digits)
	for i := 0; i < length/2; i++ {
		digits[i], digits[length-i-1] = digits[length-i-1], digits[i]
	}

	return digits
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < len(digits); j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}
