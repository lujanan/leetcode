// 2787.将一个数字表示成幂的和的方案数

package algorithm_2700

func numberOfWays(n int, x int) int {
	
}

// 回溯+剪枝
func numberOfWays1(n int, x int) int {
	var ansMap = make(map[int]int)
	var searchFn func(num, idx int) int
	searchFn = func(num, idx int) int {
		var tmp = 1
		for i := 0; i < x; i++ {
			tmp *= idx
			if tmp > num {
				return 0
			}
		}
		if tmp == num {
			return 1
		}

		if a, ok := ansMap[num|(idx<<9)]; ok {
			return a
		}

		ansMap[num|(idx<<9)] = (searchFn(num, idx+1) + searchFn(num-tmp, idx+1)) % 1000000007
		return ansMap[num|(idx<<9)]
	}

	return searchFn(n, 1)
}
