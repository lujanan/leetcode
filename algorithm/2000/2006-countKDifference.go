package algorithm_2000

func countKDifference(nums []int, k int) (res int) {
	var nMap = make(map[int]int)
	for _, v := range nums {
		if _, ok := nMap[v-k]; ok {
			res += nMap[v-k]
		}
		if _, ok := nMap[v+k]; ok {
			res += nMap[v+k]
		}
		nMap[v]++
	}
	return
}
