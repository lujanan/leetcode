package algorithm

//加一
//https://leetcode-cn.com/problems/plus-one/
func plusOne(digits []int) []int {
	var lenght = len(digits)
	for i := 0; i < lenght/2; i++ {
		digits[i], digits[lenght-i-1] = digits[lenght-i-1], digits[i]
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
	lenght = len(digits)
	for i := 0; i < lenght/2; i++ {
		digits[i], digits[lenght-i-1] = digits[lenght-i-1], digits[i]
	}

	return digits
}
