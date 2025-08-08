// 808.分汤

package algorithm_800

func soupServings(n int) float64 {
	if n > 4475 {
		return 1.00000
	}
	var resMap = make(map[int]map[int]float64)

	var soupFn func(numA, numB int) float64
	soupFn = func(numA, numB int) float64 {
		if numA <= 0 {
			if numB > 0 {
				return 1
			}
			return 0.5
		}
		if numB <= 0 {
			return 0
		}

		if res, ok := resMap[numA][numB]; ok {
			return res
		}

		var p1, p2, p3, p4 float64
		if numA <= 100 {
			p1 = 1
		} else {
			p1 = soupFn(numA-100, numB)
		}

		if numA <= 75 {
			p2 = 0.5
			if numB > 25 {
				p2 = 1
			}
		} else {
			p2 = soupFn(numA-75, numB-25)
		}

		if numA <= 50 {
			p3 = 0.5
			if numB > 50 {
				p3 = 1
			}
		} else {
			p3 = soupFn(numA-50, numB-50)
		}

		if numA <= 25 {
			p4 = 0.5
			if numB > 75 {
				p4 = 1
			}
		} else {
			p4 = soupFn(numA-25, numB-75)
		}

		p := 0.25 * (p1 + p2 + p3 + p4)
		if _, ok := resMap[numA]; !ok {
			resMap[numA] = make(map[int]float64)
		}
		resMap[numA][numB] = p
		return p
	}

	return soupFn(n, n)
}
