package algorithm_competition

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	var digMap = make(map[[2]int]int)
	for i := range dig {
		digMap[[2]int{dig[i][0], dig[i][1]}] = 1
	}

	var res int
	var canDig bool
	for _, v := range artifacts {
		y1, x1, y2, x2 := v[0], v[1], v[2], v[3]
		if _, ok := digMap[[2]int{y1, x1}]; !ok {
			continue
		}
		if _, ok := digMap[[2]int{y2, x2}]; !ok {
			continue
		}

		canDig = true
		for x1 < x2 && canDig {
			x1++
			if _, ok := digMap[[2]int{y1, x1}]; !ok {
				canDig = false
			}
		}

		x1 = v[1]
		for y1 < y2 && canDig {
			y1++
			if _, ok := digMap[[2]int{y1, x1}]; !ok {
				canDig = false
			}
		}

		if canDig {
			res++
		}
	}
	return res
}
