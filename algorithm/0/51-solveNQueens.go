package algorithm_0

import "strings"

func solveNQueens(n int) [][]string {
	var res [][]int
	var col, pie, na = make(map[int]int), make(map[int]int), make(map[int]int)
	var fn func(row int, sub []int)
	fn = func(row int, sub []int) {
		if row >= n {
			res = append(res, sub)
			return
		}

		for i := 0; i < n; i++ {
			_, okc := col[i]
			_, okx := pie[i+row]
			_, oky := na[row-i]
			if okc || okx || oky {
				continue
			}

			col[i] = 0
			pie[i+row] = 0
			na[row-i] = 0
			fn(row+1, append(append([]int{}, sub...), i))

			delete(col, i)
			delete(pie, i+row)
			delete(na, row-i)
		}
	}
	fn(0, nil)
	return format(res, n)
}

func format(nums [][]int, n int) [][]string {
	var res [][]string
	for _, num := range nums {
		var tmp []string
		for _, v := range num {
			tmp = append(tmp, strings.Repeat(".", v)+"Q"+strings.Repeat(".", n-v-1))
		}
		res = append(res, tmp)
	}
	return res
}
