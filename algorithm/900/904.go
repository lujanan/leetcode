// 904.水果成篮

package algorithm_900

func totalFruit(fruits []int) int {
	var ans, i, j int
	var numMap = make(map[int]int)
	for i < len(fruits) && j < len(fruits) {
		for ; j < len(fruits); j++ {
			numMap[fruits[j]]++
			if len(numMap) > 2 {
				break
			}
		}
		ans = max(ans, j-i)

		for ; len(numMap) > 2 && i <= j; i++ {
			numMap[fruits[i]]--
			if numMap[fruits[i]] <= 0 {
				delete(numMap, fruits[i])
			}
		}
	}
	return ans
}
