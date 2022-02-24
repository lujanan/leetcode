package algorithm_100

// 119. 杨辉三角 II
// https://leetcode-cn.com/problems/pascals-triangle-ii/

func getRow(rowIndex int) []int {
	var res = []int{1}
	for i := 1; i <= rowIndex; i++ {
		var tmp = append([]int{}, res...)
		for j := 1; j < len(res); j++ {
			res[j] += tmp[j-1]
		}
		res = append(res, 1)
	}
	return res
}
