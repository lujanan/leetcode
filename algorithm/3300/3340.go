package algorithm_3300

func isBalanced(num string) bool {
	var res int
	for i := 0; i < len(num); i++ {
		if i&1 == 0 {
			res += int(num[i] - '0')
		} else {
			res -= int(num[i] - '0')
		}
	}

	return res == 0
}
