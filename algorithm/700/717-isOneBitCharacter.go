package algorithm_700

func isOneBitCharacter(bits []int) bool {
	var one int
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 0 {
			break
		}
		one++
	}
	return one%2 == 0
}
