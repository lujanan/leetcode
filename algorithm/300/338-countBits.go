package algorithm_300

// 338. 比特位计数
// https://leetcode-cn.com/problems/counting-bits/

func countBits(n int) []int {
	var res = []int{0}
	for i := 1; i <= n; i++ {
		res = append(res, res[i>>1]+i&1)
	}
	return res
}

func countBits1(n int) []int {
	var res = []int{0}
	var j, k int
	for i := 1; i <= n; i++ {
		j, k = i, 0
		for j > 0 {
			if j&1 == 1 {
				k++
			}
			j >>= 1
		}
		res = append(res, k)
	}
	return res
}
