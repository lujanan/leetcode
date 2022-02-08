package algorithm_0

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	res := myPow(x, n>>1)
	if n%2 == 1 {
		return res * res * x
	} else {
		return res * res
	}
}
