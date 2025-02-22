// 2506.统计相似字符串对的数目

package algorithm_2500

func similarPairs(words []string) int {
	var res int
	var numMap = make(map[int]int)
	for _, w := range words {
		var num = 0
		for i := 0; i < len(w); i++ {
			num |= 1 << (w[i] - 'a')
		}
		res += numMap[num] // 直接阶加 (3组相似可组成 2+1 对，4组可组 3+2+1 对，5组可组 4+3+2+1 对)
		numMap[num]++
	}
	return res
}
