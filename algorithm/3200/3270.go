// 3270. 求出数字答案

package algorithm_3200

func generateKey(num1 int, num2 int, num3 int) int {
	var min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var num int
	for i := 1000; i > 0; i /= 10 {
		num = num*10 + min(num1/i, min(num2/i, num3/i))
		num1 %= i
		num2 %= i
		num3 %= i
	}
	return num
}
