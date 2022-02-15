package algorithm_1300

// https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/

func luckyNumbers(matrix [][]int) []int {
	var rMin = make(map[int]int)
	var cMax = make(map[int]int)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if _, ok := rMin[i]; !ok || (ok && matrix[i][rMin[i]] > matrix[i][j]) {
				rMin[i] = j
			}

			if _, ok := cMax[j]; !ok || (ok && matrix[cMax[j]][j] < matrix[i][j]) {
				cMax[j] = i
			}
		}
	}

	var res []int
	for i, j := range rMin {
		if _, ok := cMax[j]; ok && cMax[j] == i {
			res = append(res, matrix[i][j])
		}
	}
	return res
}
