package algorithm_1400

// 先以0下标分区计算，然后逐个移动计算
func maxScore(s string) (max int) {
	if len(s) < 2 {
		return 0
	}

	if s[0] == '0' {
		max++
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			max++
		}
	}
	var tmp = max
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			tmp++
		} else {
			tmp--
		}
		if tmp > max {
			max = tmp
		}
	}
	return
}
