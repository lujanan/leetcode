// 128.最长连续序列

package algorithm_100

func longestConsecutive(nums []int) int {
	var res int
	var maxMap, minMap = make(map[int]int), make(map[int]int)

	for _, num := range nums {
		maxMap[num] = maxMap[num-1] + 1
		res = max(res, maxMap[num])
		
		minMap[num] = minMap[num+1] + 1
		res = max(res, minMap[num])
	}
	return res
}
