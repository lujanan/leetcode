// 128.最长连续序列

package algorithm_100

func longestConsecutive(nums []int) int {
	var numMap = make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	
	// 1.不存在num-1,则说明num是连续序列的第一个数，开始内循环求序列长度
	// 2.否则跳过
	// 所以外循环时间复杂度是O(n)，内循环每个数也只会遍历1次，复杂度也是O(n)
	// 总的复杂度为O(n)
	var ans, curLen, curNum int
	for num := range numMap {
		if !numMap[num-1] {
			curLen = 1
			curNum = num

			for numMap[curNum+1] {
				curLen++
				curNum++
			}

			ans = max(ans, curLen)
		}
	}

	return ans
}
