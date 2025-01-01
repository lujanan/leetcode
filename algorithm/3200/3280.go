// 3280.将日期转换为二进制表示

package algorithm_3200

import "fmt"

func convertDateToBinary(date string) string {
	var res string
	var num int

	var num2B = func(num int) string {
		var str string
		for num > 0 {
			if num&1 == 1 {
				str = "1" + str
			} else {
				str = "0" + str
			}
			num = num >> 1
		}
		return str
	}

	for i := 0; i < len(date); i++ {
		if date[i] == '-' {
			if len(res) == 0 {
				res = num2B(num)
			} else {
				res += "-" + num2B(num)
			}
			num = 0
			fmt.Println(res)
		} else {
			num = num*10 + int(date[i]-'0')
		}
		fmt.Println(num)
	}
	fmt.Println(res)
	res += "-" + num2B(num)
	return res
}
