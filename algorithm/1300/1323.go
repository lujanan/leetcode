// 1323.6和9组成的最大数字

package algorithm_1300

import "math"

func maximum69Number(num int) int {
	var tmp, add = num, -1
	for i := 0; tmp > 0; i, tmp = i+1, tmp/10 {
		if tmp%10 == 6 {
			add = i
		}
	}

	if add >= 0 {
		return num + 3*int(math.Pow10(add))
	}
	return num
}
