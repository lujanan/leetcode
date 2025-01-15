// 150. 逆波兰表达式求值

package algorithm_100

import "strconv"

func evalRPN(tokens []string) int {
	var nums []int
	for _, v := range tokens {
		switch v {
		case "+":
			nums[len(nums)-2] += nums[len(nums)-1]
			nums = nums[:len(nums)-1]

		case "-":
			nums[len(nums)-2] -= nums[len(nums)-1]
			nums = nums[:len(nums)-1]

		case "*":
			nums[len(nums)-2] *= nums[len(nums)-1]
			nums = nums[:len(nums)-1]

		case "/":
			nums[len(nums)-2] /= nums[len(nums)-1]
			nums = nums[:len(nums)-1]

		default:
			num, _ := strconv.ParseInt(v, 10, 64)
			nums = append(nums, int(num))
		}
	}
	return nums[0]
}
